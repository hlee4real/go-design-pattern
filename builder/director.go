package main

type Director struct {
	builder InterfaceBuilder
}

func NewDirector(builder InterfaceBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) SetBuilder(builder InterfaceBuilder) {
	d.builder = builder
}

func (d *Director) BuildCar() Car {
	d.builder.setDoorType()
	d.builder.setColorType()
	d.builder.setWheelsType()
	return d.builder.getCar()
}
