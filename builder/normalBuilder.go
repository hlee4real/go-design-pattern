package main

type NormalBuilder struct {
	doorType   string
	colorType  string
	wheelsType string
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "Normal Basic Door"
}

func (b *NormalBuilder) setColorType() {
	b.colorType = "Normal Red Color"
}

func (b *NormalBuilder) setWheelsType() {
	b.wheelsType = "Normal 4 Wheels"
}

func (b *NormalBuilder) getCar() Car {
	return Car{
		doorType:   b.doorType,
		colorType:  b.colorType,
		wheelsType: b.wheelsType,
	}
}
