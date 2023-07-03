package main

import "fmt"

//visitor la 1 DP cho phep them nhung hanh dong moi cho cac class da ton tai nhung khong thay doi chung no

//trong truong hop nay, ta co 3 class: square, circle, rectangle
//ta muon them 2 hanh dong: tinh dien tich va tinh toa do trung tam
//ta co the them 2 ham tinh dien tich va tinh toa do trung tam vao cac class nay
//nhung neu ta muon them 1 hanh dong nua thi ta phai them 1 ham nua vao cac class nay
//vi vay ta su dung visitor de them hanh dong nay ma khong can thay doi cac class da ton tai

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
