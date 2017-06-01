package server

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type Message struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

func (self *Message) String() string {
	return self.Author + " says " + self.Body
}