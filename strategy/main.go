package main

// strategy la mot DP co the thay doi mot so luong hanh vi thanh cac doi tuong rieng biet va xoay vong chung

func main() {
	paper := &Paper{}
	rock := &Rock{}

	game := initGame(paper)
	game.choose(rock)
}
