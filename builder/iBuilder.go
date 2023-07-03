package main

//Builder la mot design pattern thuoc creational patter, duoc su dung khi san pham tao ra phuc tap
//va can nhieu buoc de hoan thanh. trong Builder pattern, chung ta tach qua trinh xay dung mot object
//phuc tap ra khoi qua trinh tao ra no, de chung ta co the tao ra nhieu object phuc tap tuong tu nhau
//nhung qua trinh xay dung khac nhau.

type InterfaceBuilder interface {
	setDoorType()
	setColorType()
	setWheelsType()
	getCar() Car
}

func getBuilder(builderType string) InterfaceBuilder {
	if builderType == "Normal" {
		return newNormalBuilder()
	}

	if builderType == "Sports" {
		return newSportsBuilder()
	}
	return nil
}
