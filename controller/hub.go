package controller

import (
	"log"
	"net/http"

	"github.com/mayur-tolexo/chat/model"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const (
	socketBufferSize = 2048
)

var (
	upgrader   = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}
	defaultHub = model.NewHub()
)

// HubHandler of websocket
func HubHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)

		socket, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			log.Println("Upgrade error", err)
			return
		}

		// // Get user out of session
		// user, ok := SessUserDetail(req)
		// if !ok {
		// 	log.Println("Invaid user session")
		// 	return
		// }

		channel := vars["id"]
		client := model.NewClient(socket, defaultHub, channel)
		// client.SetUser(user)
		// client.SetSave(model.NewSaveMessageChan())

		// //checking user permission to access the channel
		// if err := core.ChanPermCheck(user, channel); err != nil {
		// 	log.Println("Not Authorised")
		// 	return
		// }

		defaultHub.Join(client)
		defer func() { defaultHub.Leave(client) }()
		// go client.Write()
		// client.Read()
	}
}
