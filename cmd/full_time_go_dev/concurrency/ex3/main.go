package main

import (
	"fmt"
)

type Server struct {
	quitch chan struct{}
	msgch  chan string
}

func NewServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string),
	}
}

func (s *Server) Start() {
	fmt.Println("Server starting")
	s.Loop()
}

func (s *Server) Loop() {
outerloop:
	for {
		select {
		case <-s.quitch:
			break outerloop
		case msg := <-s.msgch:
			s.handleMessage(msg)
		default:
			// waiting for new messages
		}
	}
}

func (s *Server) quit() {
	s.quitch <- struct{}{}
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("We received a message: " + msg)
}

func (s *Server) sendMessage(msg string) {
	s.msgch <- msg
}

func main() {
	server := NewServer()
	go server.Start()

	for i := 0; i < 100; i++ {
		server.sendMessage(fmt.Sprintf("Handle: %d", i))
	}
}
