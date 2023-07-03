package main

import "fmt"

type AppStore struct {
}

func (a *AppStore) SendMessage() {
	fmt.Println("Sending Message from App Store")
}
