package main

import "fmt"

type Rock struct {
}

func (r *Rock) choose(g *Game) {

	fmt.Println("Rock")
}
