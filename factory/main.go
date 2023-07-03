package main

import "fmt"

// dinh nghia: la 1 design pattern thuoc creational design,
// cung cap interfaces cho viec tao ra cac doi tuong o lop cao, nhung cho phep cac lop con thay doi kieu doi tuong duoc tao ra

func main() {
	xeDapNhatBan, err := getBike("XeDapNhatBan")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(xeDapNhatBan.getName())

}
