package main

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/ibrahimker/golang-praisindo-advanced/session-21-latihan-cqrs/config"
	"github.com/ibrahimker/golang-praisindo-advanced/session-21-latihan-cqrs/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"os"
	"time"
)

var writeDB *gorm.DB

type writeDBInserterConsumerGroupHandler struct{}

func (writeDBInserterConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (writeDBInserterConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h writeDBInserterConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		slog.Info("Received message", slog.Any("string(msg.Value)", string(msg.Value)))
		var user entity.User
		_ = json.Unmarshal(msg.Value, &user)
		slog.Info("Unmarshall results", slog.Any("user", user))
		query := "INSERT INTO users (id, name) VALUES ($1, $2)"
		if err := writeDB.Exec(query, user.ID, user.Name).Error; err != nil {
			log.Printf("Error deleting user: %v\n", err)
			return err
		}
		// Process the message as per your requirement here
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	var err error
	writeDB, err = gorm.Open(postgres.Open(config.DBWriteDSN), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		slog.Error("error opening write database", slog.String("dsn", config.DBWriteDSN), slog.Any("err", err))
		os.Exit(1)
	}

	brokers := []string{config.KafkaBrokerAddress}
	groupID := config.WriteDBInserterConsumerGroupID
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V3_6_0_0 // specify appropriate Kafka version
	saramaConfig.Consumer.Offsets.AutoCommit.Enable = true
	saramaConfig.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, saramaConfig)
	if err != nil {
		slog.Error("error when setup sarama.NewConsumerGroup:", slog.Any("error", err))
		os.Exit(1)
	}

	ctx := context.Background()
	slog.InfoContext(ctx, "Start consuming from topic", slog.Any("topic", config.TopicInsertUser))
	for {
		if err := consumerGroup.Consume(ctx, []string{config.TopicInsertUser}, writeDBInserterConsumerGroupHandler{}); err != nil {
			slog.Error("error when call consumerGroup.Consume:", slog.Any("error", err))
		}
	}

}
