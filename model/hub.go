package model

// Hub represents a websocket hub having multiple clients per room
type Hub struct {
	// broadcast will broadcast the message to all the clients connect to same room
	broadcast chan *Message
	// join is a channel to add new client to a room
	join chan *Client
	// leave is a channel to remove a client from a room.
	leave chan *Client
	// clients contains all the client belongs to same room
	clients map[string]map[*Client]struct{}
}

// Join will forward the client to a join
func (h *Hub) Join(client *Client) {
	h.join <- client
}

// Leave will forward the client to remove
func (h *Hub) Leave(client *Client) {
	h.leave <- client
}

// run start a room and run it forever
func run(h *Hub) {
	for {
		select {
		case client := <-h.join:
			// joining
			if _, exists := h.clients[client.room]; !exists {
				h.clients[client.room] = make(map[*Client]struct{})
			}
			h.clients[client.room][client] = struct{}{}
		case client := <-h.leave:
			// leaving
			delete(h.clients[client.room], client)
		case msg := <-h.broadcast:
			// boradcast message to all clients
			for client := range h.clients[msg.Room] {
				client.send <- msg
			}
		}
	}
}

// NewHub creates a new websocket hub where clients will join in respective channel
func NewHub() *Hub {
	r := &Hub{
		broadcast: make(chan *Message),
		join:      make(chan *Client),
		leave:     make(chan *Client),
		clients:   make(map[string]map[*Client]struct{}),
	}
	go run(r)
	return r
}
