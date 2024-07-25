package evhandler

import (
	"context"
	"encoding/json"
	kafkaPkg "mysql-binlog/kafka"
	"strconv"
	"time"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/patrickmn/go-cache"
	"github.com/segmentio/kafka-go"
	"github.com/siddontang/go-log/log"
)

const (
	MIN_BATCH_SIZE = 5000
)

type MyEventHandler struct {
	canal.DummyEventHandler
	waitingQueue Queue
	holding      *cache.Cache
	producer     kafkaPkg.Producer
	msgCh        chan kafka.Message
	releaseCh    chan<- CommitLogInfo
}

func NewMyEventHandler(producer kafkaPkg.Producer, releaseCh chan<- CommitLogInfo) *MyEventHandler {
	msgCh := make(chan kafka.Message, 100000)
	cache := cache.New(5*time.Minute, 10*time.Minute)

	return &MyEventHandler{
		waitingQueue: Queue{},
		holding:      cache,
		producer:     producer,
		msgCh:        msgCh,
		releaseCh: releaseCh,
	}
}

// Reader Thread
func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	h.waitingQueue.Enqueue(e.Header.LogPos)

	payload, err := json.Marshal(e.Rows)
	if err != nil {
		return err
	}

	h.msgCh <- kafka.Message{
		Topic: e.Table.Name,
		Value: payload,
		Headers: []kafka.Header{
			{
				Key:   "server-id",
				Value: []byte(strconv.Itoa(int(e.Header.LogPos))),
			},
			{
				Key:   "action",
				Value: []byte(e.Action),
			},
			{
				Key:   "log-pos",
				Value: []byte(strconv.Itoa(int(e.Header.LogPos))),
			},
		},
	}

	return nil
}

// Pending queue and Kafka producer thread
func (h *MyEventHandler) StartProduceMessage(ctx context.Context) {
	ticker := time.NewTicker(100*time.Millisecond)
	defer ticker.Stop()

	var msgs []kafka.Message
	for {
		select {
		case info, ok := <-h.msgCh:
			if !ok {
				// Channel is closed and drained
				if len(msgs) > 0 {
					err := h.producer.PublishMessage(ctx, msgs...)
					if err != nil {
						log.Infof("Error when publish message: %v", err)
					}
				}
				return
			}
			msgs = append(msgs, info)

		case <-ticker.C:
			if len(msgs) > MIN_BATCH_SIZE {
				log.Infof("Published %d messages\n", len(msgs))
				err := h.producer.PublishMessage(ctx, msgs...)
				if err != nil {
					log.Infof("Error when publish message: %v", err)
				}

				for _, msg := range msgs {
					pos, _ := strconv.ParseUint(string(msg.Headers[2].Value), 10, 32)
					h.holding.Set(string(msg.Headers[2].Value), CommitLogInfo{
						Position: uint32(pos),
					}, 5*time.Minute)
				}
				msgs = nil
			}
		}
	}
}

func (h *MyEventHandler) Watcher() {
	for {
		key, ok := h.waitingQueue.Peek()
		if !ok {
			// Queue is empty, sleep for a while before checking again
			time.Sleep(100 * time.Millisecond)
			continue
		}

		cacheKey := strconv.Itoa(int(key))
		value, exists := h.holding.Get(cacheKey)

		if exists {
			// Dequeue the item since the condition is met
			_, _ = h.waitingQueue.Dequeue()
			h.releaseCh <- value.(CommitLogInfo)
			h.holding.Delete(cacheKey)
		} else {
			// Wait before checking again
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}
