package main

type InterfaceShirt interface {
	setBrand(name string)
	getBrand() string
}

type Shirt struct {
	Brand string
}

type PremierLeagueShirt struct {
	Shirt
}

type ChampionsLeagueShirt struct {
	Shirt
}

func (s *Shirt) setBrand(name string) {
	s.Brand = name
}

func (s *Shirt) getBrand() string {
	return s.Brand
}
