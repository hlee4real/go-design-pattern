package main

import "fmt"

type Game struct {
	lastMove   ChooseAlgo
	chooseAlgo ChooseAlgo
}

func initGame(c ChooseAlgo) *Game {
	return &Game{chooseAlgo: c}
}

func (g *Game) choose(c ChooseAlgo) {
	paper := &Paper{}
	rock := &Rock{}
	scissors := &Scissors{}
	if g.lastMove == nil {
		g.lastMove = c
		g.chooseAlgo.choose(g)
	}

	if g.lastMove == c {
		fmt.Println("draw")
		return
	}

	if g.lastMove == paper && c == rock {
		fmt.Println("paper wins")
		return
	}
	if g.lastMove == paper && c == scissors {
		fmt.Println("scissors wins")
		return
	}

	if g.lastMove == rock && c == paper {
		fmt.Println("paper wins")
		return
	}

	if g.lastMove == rock && c == scissors {
		fmt.Println("rock wins")
		return
	}

}
