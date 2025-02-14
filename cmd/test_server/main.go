package main

import (
	"context"
	"encoding/json"
	"github.com/VikaPaz/message_server/internal/client/queue"
	"github.com/VikaPaz/message_server/internal/models"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"time"
)

type MassageRead struct {
	ID      string `json:"ID"`
	Massage string `json:"Massage"`
}

type MessageWrite struct {
	ID     string
	Status models.Status
}

func main() {
	if err := godotenv.Overload("env/.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	confWrite := queue.Config{
		Topic:     "topic2",
		Partition: 0,
		Host:      os.Getenv("KAFKA_ADDRESS"),
		Network:   "tcp",
	}
	writer, err := queue.Connection(confWrite)
	if err != nil {
		log.Fatalf("Error connecting to kafka: %v", err)
	}

	confRead := kafka.ReaderConfig{
		Topic:     "topic1",
		Partition: 0,
		GroupID:   "g1",
		MaxWait:   24 * time.Hour * 10,
		Brokers:   []string{os.Getenv("KAFKA_ADDRESS")},
	}

	reader := kafka.NewReader(confRead)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %s\n", err)
			continue
		}
		log.Println("Message: ", string(msg.Value))

		r := MassageRead{}
		err = json.Unmarshal(msg.Value, &r)
		if err != nil {
			log.Printf("Error decoding message: %s\n", err)
			continue
		}

		w := MessageWrite{
			ID:     r.ID,
			Status: models.StatusCompleted,
		}

		value, err := json.Marshal(w)
		if err != nil {
			log.Println("failed to serialize structure: %v", err)
			continue
		}

		_, err = writer.WriteMessages(
			kafka.Message{Value: value},
		)
		if err != nil {
			log.Println("failed to write messages: %v", err)
			continue
		}
	}
}
