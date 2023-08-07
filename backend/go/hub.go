package main

type WebsocketHandler interface {
	HandleMessage(msg []byte) []byte
}

type Hub struct {
	clients map[*Client]bool

	register chan *Client

	unregister chan *Client

	broadcast chan []byte

	handler WebsocketHandler
}

func newHub(handler WebsocketHandler) *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
		handler:    handler,
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case msg := <-h.broadcast:
			res := h.handler.HandleMessage(msg)
			for client := range h.clients {
				select {
				case client.send <- res:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
