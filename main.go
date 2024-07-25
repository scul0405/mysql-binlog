package main

import (
	"context"
	"log"
	evhandler "mysql-binlog/ev_handler"
	testdata "mysql-binlog/test_data"
	"os"
	"strings"
	"time"

	kafkaPkg "mysql-binlog/kafka"

	"github.com/go-mysql-org/go-mysql/canal"
)

func main() {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = os.Getenv("MYSQL_ADDR")
	cfg.User = "root"
	cfg.Password = os.Getenv("MYSQL_ROOT_PASSWORD")
	cfg.Dump.TableDB = os.Getenv("MYSQL_DATABASE")
	cfg.Dump.Tables = strings.Split(os.Getenv("TABLES"), ",")

	c, err := canal.NewCanal(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Init
	producer := kafkaPkg.NewProducer([]string{"host.docker.internal:9092"})
	info := make(chan evhandler.CommitLogInfo, 10000)
	commitLog := evhandler.NewCommitLog(producer)

	// Register a handler to handle RowsEvent
	myEventHandler := evhandler.NewMyEventHandler(producer, info)
	c.SetEventHandler(myEventHandler)

	// Start
	ctx := context.Background()
	go myEventHandler.StartProduceMessage(ctx)
	go myEventHandler.Watcher()
	go commitLog.ReleaseCheckpoint(ctx, info)

	go func() {
		time.Sleep(5 * time.Second)

		testdata.Insert()
	}()

	c.Run()
}
