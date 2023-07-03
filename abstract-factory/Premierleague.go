package main

type PremierLeague struct{}

func (p PremierLeague) CreateBall() InterfaceBall {
	return &Ball{
		Brand: "Premier League",
	}
}

func (p PremierLeague) CreateShirt() InterfaceShirt {
	return &Shirt{
		Brand: "Premier League",
	}
}
