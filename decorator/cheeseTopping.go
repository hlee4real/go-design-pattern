package main

type CheeseTopping struct {
	milktea InterfaceMilktea
}

func (c *CheeseTopping) getPrice() int {
	return c.milktea.getPrice() + 10000
}
