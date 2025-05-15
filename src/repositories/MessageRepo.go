package repositories

import (
	"errors"
	"fmt"
	"jwt-authentication/database"
	"jwt-authentication/models"
)

type messageRepository struct{}

func NewMessageRepository() MessageRepository {
	return &messageRepository{}
}

func (m *messageRepository) Save(msg *models.Message) error {
	stmt := database.SqlMessageDb("saveMessage")

	result, err := database.DB.Exec(stmt, msg.Sender, msg.Receiver, msg.Content, msg.Timestamp)
	if err != nil {
		fmt.Println(err, "database err")
		return errors.New("error saving message")
	}

	// Adds id from db to json data
	msgId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	msg.Id = int(msgId)

	return nil
}
func (m *messageRepository) GetMsg(sender, receiver string) ([]models.Message, error) {
	stmt := database.SqlMessageDb("getMessage")
	rows, err := database.DB.Query(stmt, sender, receiver, receiver, sender)
	if err != nil {
		return nil, err
	}

	var messages []models.Message
	for rows.Next() {
		var message models.Message
		err := rows.Scan(&message.Sender, &message.Receiver, &message.Content, &message.Timestamp)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
