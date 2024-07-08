package main

import (
	"fmt"

	//_ "github.com/joho/godotenv/autoload"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	kafkaBroker = "kafka:9092"
	topic       = "ip"
	groupID     = "ip-group"
)

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaBroker,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		return
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Printf("Failed to subscribe to topics: %s\n", err)
		return
	}

	db, err := db()
	if err != nil {
		fmt.Printf("Failed to connect to database: %s\n", err)
		return
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			result := db.Create(&IP{IP: string(msg.Value)})
			if result.Error != nil {
				fmt.Printf("Failed to insert message: %s\n", result.Error)
				return
			}
			fmt.Printf("Received message: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v\n", err)
		}
	}
}

type IP struct {
	gorm.Model
	IP string `gorm:"not null"`
}

func db() (*gorm.DB, error) {
	dsn := "root:@tcp(mariadb:3306)/logger?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if !db.Migrator().HasTable(&IP{}) {
		if err := db.AutoMigrate(&IP{}); err != nil {
			return nil, err
		}
	}

	return db, nil
}
