package main

import (
	"context"
	"github.com/hrittikhere/kafka-demo/faker"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

const (
	topic         = "quickstart-events"
	brokerAddress = "localhost:9092"
	partition     = 0
)

func main() {

	conn, err := kafka.DialLeader(context.Background(), "tcp", brokerAddress, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(cmd.GetFakeUser())},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

}
