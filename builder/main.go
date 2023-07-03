package main

import "fmt"

// Builder la mot design pattern thuoc creational patter, duoc su dung khi san pham tao ra phuc tap
// va can nhieu buoc de hoan thanh. trong Builder pattern, chung ta tach qua trinh xay dung mot object
// phuc tap ra khoi qua trinh tao ra no, de chung ta co the tao ra nhieu object phuc tap tuong tu nhau
// nhung qua trinh xay dung khac nhau.
func main() {
	normalBuilder := getBuilder("Normal")
	sportBuilder := getBuilder("Sports")

	director := NewDirector(normalBuilder)
	normalCar := director.BuildCar()

	fmt.Println(normalCar.doorType)
	fmt.Println(normalCar.colorType)
	fmt.Println(normalCar.wheelsType)

	director.SetBuilder(sportBuilder)
	sportCar := director.BuildCar()

	fmt.Println(sportCar.doorType)
	fmt.Println(sportCar.colorType)
	fmt.Println(sportCar.wheelsType)
}
