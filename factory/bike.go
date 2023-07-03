package main

// dinh nghia: la 1 design pattern thuoc creational design,
// cung cap interfaces cho viec tao ra cac doi tuong o lop cao, nhung cho phep cac lop con thay doi kieu doi tuong duoc tao ra

type InterfaceVehicle interface {
	setName(name string)
	setSpeed(speed int)
	getName() string
	getSpeed() int
}

type Bike struct {
	name  string
	speed int
}

func (b *Bike) setName(name string) {
	b.name = name
}

func (b *Bike) setSpeed(speed int) {
	b.speed = speed
}

func (b *Bike) getName() string {
	return b.name
}

func (b *Bike) getSpeed() int {
	return b.speed

}
