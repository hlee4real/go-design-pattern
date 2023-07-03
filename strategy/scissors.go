package main

import "fmt"

type Scissors struct {
}

func (s *Scissors) choose(g *Game) {
	fmt.Println("Scissors")
}
