package main

import (
	"fmt"
	"log"
)

// facade la structural pattern cung cap nhung interfaces don gian cho nhung he thong phuc tap
// cua cac lop,thu vien hoac framework

//trong truong hop nay, chung ta se tao ra mot facade de cung cap nhung interfaces don gian
// cho cac he thong phuc tap cua cac lop trong mot he thong thanh toan
// trong he thong thanh toan nay, chung ta co cac lop sau:

// 1. Ledger: lop nay se luu tru cac giao dich cua nguoi dung
// 2. SecurityCode: lop nay se kiem tra ma bao mat cua nguoi dung
// 3. Notification: lop nay se gui cac thong bao cho nguoi dung
// 4. Wallet: lop nay se luu tru so tien cua nguoi dung
// 5. Account: lop nay se kiem tra tai khoan cua nguoi dung
// 6. WalletFacade: lop nay se cung cap cac interfaces don gian cho cac lop phuc tap

func main() {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
