package main

import "fmt"

type IPhone struct {
	messenger Messenger
}

func (i *IPhone) Message() {
	fmt.Println("Message request for iphone")
	i.messenger.SendMessage()
}

func (i *IPhone) setMessage(m Messenger) {
	i.messenger = m
}
