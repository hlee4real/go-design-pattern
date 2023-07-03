package main

import "fmt"

type Paper struct {
}

func (p *Paper) choose(g *Game) {
	fmt.Println("Paper")
}
