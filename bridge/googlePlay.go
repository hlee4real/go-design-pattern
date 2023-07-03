package main

import "fmt"

type GooglePlay struct {
}

func (g *GooglePlay) SendMessage() {
	fmt.Println("Sending Message from Google Play")
}
