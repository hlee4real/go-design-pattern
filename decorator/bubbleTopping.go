package main

type BubbleTopping struct {
	milktea InterfaceMilktea
}

func (b *BubbleTopping) getPrice() int {
	return b.milktea.getPrice() + 5000
}
