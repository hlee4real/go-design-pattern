package main

import "fmt"

//bridge chia nho 1 class lon thanh cac class nho hon de de~ phat trien doc lap
//bridge giup tach biet abstraction va implementation

func main() {
	appstore := &AppStore{}
	googleplay := &GooglePlay{}

	iphone := &IPhone{}
	iphone.setMessage(appstore)
	iphone.Message()
	fmt.Println()

	iphone.setMessage(googleplay)
	iphone.Message()
	fmt.Println()

	samsung := &Samsung{}
	samsung.setMessage(googleplay)
	samsung.Message()
	fmt.Println()

	samsung.setMessage(appstore)
	samsung.Message()
	fmt.Println()
}
