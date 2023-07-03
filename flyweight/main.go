package main

import "fmt"

// flyweight la mot structural design pattern, giup giam dung luong bo nho su dung boi cac doi tuong
//
// flyweight pattern se chia cac doi tuong thanh 2 loai:
// - intrinsic: chua cac thong tin co the chia se giua cac doi tuong
// - extrinsic: chua cac thong tin khong the chia se giua cac doi tuong

//trong truong hop nay, chung ta se tao ra cac doi tuong dress, chia thanh 2 loai:
// - intrinsic: chua thong tin co the chia se giua cac doi tuong, trong truong hop nay la mau sac
// - extrinsic: chua thong tin khong the chia se giua cac doi tuong, trong truong hop nay la kieu cua doi tuong

func main() {
	game := newGame()

	//Add Terrorist
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	//Add CounterTerrorist
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
