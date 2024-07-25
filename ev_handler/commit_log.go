package evhandler

import (
	"context"
	"encoding/json"
	kafkaPkg "mysql-binlog/kafka"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/siddontang/go-log/log"
)

type CommitLogInfo struct {
	Position uint32
	FileName string
}

type CommitLog struct {
	producer kafkaPkg.Producer
	reader   kafka.Reader
}


// TODO: add reader here
func NewCommitLog(producer kafkaPkg.Producer) *CommitLog {
	return &CommitLog{
		producer: producer,
	}
}

// For starting
func (c *CommitLog) GetLastCommitLogInfo(ctx context.Context) CommitLogInfo {
	lastMsg, err := c.reader.FetchMessage(ctx)
	if err != nil {
		log.Infof("Error when get last commit log info: %v", err)
	}

	var info CommitLogInfo
	err = json.Unmarshal(lastMsg.Value, &info)
	if err != nil {
		log.Infof("Error when unmarshal last commit log info: %v", err)
	}

	return info
}

// Commit Checkpoint Thread
func (c *CommitLog) ReleaseCheckpoint(ctx context.Context, infoCh <-chan CommitLogInfo) {
	// Threshold to release checkpoint (batching)
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	var msgs []kafka.Message
	for {
		select {
		case info, ok := <-infoCh:
			if !ok {
				// Channel is closed and drained
				if len(msgs) > 0 {
					err := c.producer.PublishMessage(ctx, msgs...)
					if err != nil {
						log.Infof("Error when publish checkpoint: %v", err)
					}
				}
				return
			}
			jsonData, _ := json.Marshal(info)
			msg := kafka.Message{
				Topic: "commit-log",
				Value: jsonData,
			}
			msgs = append(msgs, msg)

		case <-ticker.C:
			if len(msgs) > 0 {
				err := c.producer.PublishMessage(ctx, msgs...)
				if err != nil {
					log.Infof("Error when publish checkpoint: %v", err)
				}
				msgs = nil
			}
		}
	}
}
