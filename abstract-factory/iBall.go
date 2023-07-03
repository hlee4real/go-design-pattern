package main

type InterfaceBall interface {
	setBrand(brand string)
	getBrand() string
}

type Ball struct {
	Brand string
}

type PremierLeagueBall struct {
	Ball
}

type ChampionsLeagueBall struct {
	Ball
}

func (b *Ball) setBrand(brand string) {
	b.Brand = brand
}

func (b *Ball) getBrand() string {
	return b.Brand
}
