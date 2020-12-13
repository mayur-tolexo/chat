package model

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	messageBufferSize = 1024
)

// Client represents a user connect to a room, one user may have many devices to chat,
// so it should not be the same as user
type Client struct {
	// socket is the web socket for this client.
	socket *websocket.Conn
	// room name to which client is joined
	room string
	// send is a channel on which messages are sent.
	send chan *Message
	// hub is the room this client is chatting in.
	hub *Hub
	// user is interface which contains client information
	user User
}

// NewClient created
func NewClient(socket *websocket.Conn, h *Hub, room string) *Client {
	client := &Client{
		socket: socket,
		hub:    h,
		room:   room,
		send:   make(chan *Message, messageBufferSize),
	}
	return client
}

// Read from client socket
func (c *Client) Read() {
	defer c.socket.Close()
	for {
		var msg *Message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			log.Print(err)
			return
		}

		msg.ID = uuid.New().String()
		msg.Room = c.room
		msg.LoginID = c.user.GetUserLoginID()
		msg.CreateAt = time.Now().UTC()
		c.hub.broadcast <- msg
		// // send message to save in another channel
		// *c.save <- *sm
	}
}

// Write in client socket
func (c *Client) Write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}
