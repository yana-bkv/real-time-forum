package websocket

import "github.com/gorilla/websocket"

type Client struct {
	ID     string
	PeerID string
	Conn   *websocket.Conn
	Send   chan []byte
}

type Message struct {
	From    string
	To      string
	Content []byte
}

type Hub struct {
	Clients    map[string]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
}

var HubInstance = Hub{
	Clients:    make(map[string]*Client),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Broadcast:  make(chan Message),
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.ID] = client
		case client := <-h.Unregister:
			if _, ok := h.Clients[client.ID]; ok {
				delete(h.Clients, client.ID)
				close(client.Send)
			}
		case msg := <-h.Broadcast:
			if receiver, ok := h.Clients[msg.To]; ok {
				receiver.Send <- msg.Content
			}
		}
	}
}
