package main

type SportBuilder struct {
	doorType   string
	colorType  string
	wheelsType string
}

func newSportsBuilder() *SportBuilder {
	return &SportBuilder{}
}

func (b *SportBuilder) setDoorType() {
	b.doorType = "Sport Luxury Door"
}

func (b *SportBuilder) setColorType() {
	b.colorType = "Sport Red and Blue Color"
}

func (b *SportBuilder) setWheelsType() {
	b.wheelsType = "Sport 4 spax Wheels"
}

func (b *SportBuilder) getCar() Car {
	return Car{
		doorType:   b.doorType,
		colorType:  b.colorType,
		wheelsType: b.wheelsType,
	}
}
