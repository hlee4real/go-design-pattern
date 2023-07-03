package main

type ChampionsLeague struct{}

func (c ChampionsLeague) CreateBall() InterfaceBall {
	return &Ball{
		Brand: "Champions League",
	}
}

func (c ChampionsLeague) CreateShirt() InterfaceShirt {
	return &Shirt{
		Brand: "Champions League",
	}
}
