package controllers

import (
	"github.com/gorilla/mux"
	"jwt-authentication/services"
	"net/http"
)

type MessageController struct {
	messageService services.MessageService
}

func NewMessageController(messageService services.MessageService) *MessageController {
	return &MessageController{messageService}
}

//func (c *MessageController) AddMessage(w http.ResponseWriter, r *http.Request) {
//	var data string
//
//	err := DecodeJson(r, w, &data)
//	if err != nil {
//		return
//	}
//
//	vars := mux.Vars(r)
//	userId := vars["user"]
//	peerId := vars["peer"]
//
//	//authorId, err := strconv.Atoi(helpers.GetUserId(w, r))
//	//if err != nil {
//	//	http.Error(w, "Invalid request payload for post", http.StatusBadRequest)
//	//}
//
//	msg, err := c.messageService.Save(userId, peerId, data)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	err = EncodeJson(w, msg)
//	if err != nil {
//		return
//	}
//}

func (c *MessageController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user"]
	peerId := vars["peer"]

	msgs, err := c.messageService.Get(userId, peerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Encode response as JSON
	err = EncodeJson(w, msgs)
	if err != nil {
		return
	}
}
