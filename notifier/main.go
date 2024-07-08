package main

import (
	"fmt"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	kafkaBroker = "kafka:9092"
)

var (
	topic = "ip"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaBroker})
	if err != nil {
		fmt.Fprintf(w, "Failed to create producer: %s\n", err)
		return
	}
	defer producer.Close()

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(r.RemoteAddr),
	}

	err = producer.Produce(message, nil)
	if err != nil {
		fmt.Fprintf(w, "Failed to produce message: %s\n", err)
		return
	}

	fmt.Fprintf(w, "Welcome, %s", r.RemoteAddr)
}
