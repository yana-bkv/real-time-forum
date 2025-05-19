package services

import (
	"database/sql"
	"errors"
	"fmt"
	"jwt-authentication/models"
	"jwt-authentication/repositories"
	"time"
)

type MessageServiceImpl struct {
	messageRepo repositories.MessageRepository
}

func NewMessageService(messageRepo repositories.MessageRepository) *MessageServiceImpl {
	return &MessageServiceImpl{messageRepo: messageRepo}
}

func (m *MessageServiceImpl) Save(sender, receiver, content string) (*models.Message, error) {
	timeNow := time.Now()

	msg := models.Message{
		Sender:    sender,
		Receiver:  receiver,
		Content:   content,
		Timestamp: timeNow.Format("2006-01-02 15:04:05"),
	}

	fmt.Println(msg)
	err := m.messageRepo.Save(&msg)
	if err != nil {
		return nil, errors.New("Unable to save message")
	}
	return &msg, nil
}

func (m *MessageServiceImpl) Get(sender, receiver string) ([]models.Message, error) {
	msgs, err := m.messageRepo.GetMsg(sender, receiver)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Database error messages")
		}
		return nil, err
	}
	return msgs, nil
}
