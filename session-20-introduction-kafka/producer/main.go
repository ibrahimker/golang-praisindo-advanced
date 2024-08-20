package main

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/ibrahimker/golang-praisindo-advanced/session-20-introduction-kafka/entity"
	"log/slog"
	"os"
	"time"
)

func main() {
	brokers := []string{entity.KafkaBrokerAddress}
	producer, err := sarama.NewSyncProducer(brokers, nil)
	if err != nil {
		slog.Error("error when call sarama.NewSyncProducer", slog.Any("error", err))
		os.Exit(1)
	}

	message := entity.Event{Message: "Hello world at " + time.Now().Format(time.DateTime)}
	jsonMessage, _ := json.Marshal(message)
	msg := &sarama.ProducerMessage{
		Topic: entity.DefaultTopic,
		Value: sarama.StringEncoder(jsonMessage),
	}
	_, _, err = producer.SendMessage(msg)
	if err != nil {
		slog.Error("error when call producer.SendMessage ", slog.Any("error", err))
	}
	slog.Info("Successfully sent message", slog.Any("message", message))
}
