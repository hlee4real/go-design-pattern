package main

import "fmt"

type Samsung struct {
	messenger Messenger
}

func (s *Samsung) Message() {
	fmt.Println("Message request for samsung")
	s.messenger.SendMessage()
}

func (s *Samsung) setMessage(m Messenger) {
	s.messenger = m
}
