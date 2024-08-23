package handler

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/session-21-latihan-cqrs/config"
	"github.com/ibrahimker/golang-praisindo-advanced/session-21-latihan-cqrs/entity"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

type UserHandler struct {
	readDB   *gorm.DB
	producer sarama.SyncProducer
}

func NewUserHandler(readDB *gorm.DB, producer sarama.SyncProducer) *UserHandler {
	return &UserHandler{readDB: readDB, producer: producer}
}

// GetAllUsers returns all user from read DB
func (u *UserHandler) GetAllUsers(c *gin.Context) {
	var users []entity.User
	if err := u.readDB.Raw("select id,name from users").Scan(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser publish user data to topic config.TopicInsertUser
func (u *UserHandler) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonMessage, _ := json.Marshal(user)
	msg := &sarama.ProducerMessage{
		Topic: config.TopicInsertUser,
		Value: sarama.StringEncoder(jsonMessage),
	}

	if _, _, err := u.producer.SendMessage(msg); err != nil {
		slog.Error("error when call producer.SendMessage ", slog.Any("error", err))
	}
	slog.Info("Successfully sent user", slog.Any("user", user))
	c.JSON(http.StatusCreated, user)
}
