package main

import (
	"fmt"
)

// state pattern cho phep mot doi tuong thay doi hanh vi cua no khi trang thai noi tai cua no thay doi
// trong truong hop nay, chung ta se thay doi hanh vi cua mot may ban hang tuy thuoc vao trang thai cua no
// khi may ban hang co hang, no se cho phep nguoi dung chon hang
// khi nguoi dung chon hang, no se yeu cau nguoi dung tra tien
// khi nguoi dung tra tien, no se ban hang
// khi may ban hang het hang, no se khong cho phep nguoi dung chon hang
// khi may ban hang het hang, no se khong cho phep nguoi dung tra tien
// khi may ban hang het hang, no se khong ban hang
// khi may ban hang het hang, no se cho phep nguoi dung them hang
// ...
func main() {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()

	err = vendingMachine.addItem(2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()

	err = vendingMachine.requestItem()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		fmt.Println(err)
		return
	}
}
