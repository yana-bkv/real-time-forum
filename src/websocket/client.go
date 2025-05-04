package websocket

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user"]
	peerID := vars["peer"]

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	client := &Client{
		ID:     userID,
		PeerID: peerID,
		Conn:   conn,
		Send:   make(chan []byte, 256),
	}

	HubInstance.Register <- client

	go client.writePump()
	go client.readPump()
}

func (c *Client) readPump() {
	defer func() {
		// Закрытие соединения и удаление клиента из хаба
		if err := recover(); err != nil {
			log.Printf("Error during readPump: %v", err)
		}
		HubInstance.Unregister <- c
		c.Conn.Close()
	}()

	for {
		// Чтение сообщения
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			// Логируем ошибку при чтении
			log.Printf("Error reading message from client %s: %v", c.ID, err)
			break
		}

		// Отправка сообщения в Broadcast канал
		err = sendMessageToBroadcast(c, message)
		if err != nil {
			log.Printf("Error broadcasting message from %s: %v", c.ID, err)
			break
		}
	}
}

// Функция для отправки сообщения в Broadcast канал с обработкой ошибок
func sendMessageToBroadcast(c *Client, message []byte) error {
	// Проверка на наличие получателя
	if receiver, ok := HubInstance.Clients[c.PeerID]; ok {
		receiver.Send <- message
		log.Printf("Message from %s to %s: %s", c.ID, c.PeerID, string(message))
	} else {
		log.Printf("Receiver %s is offline", c.PeerID)
		//return fmt.Errorf("receiver not found")
	}

	return nil
}

func (c *Client) writePump() {
	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
	c.Conn.Close()
}
