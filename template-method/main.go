package main

import "fmt"

// template-method la mot design pattern cho phep chung ta dinh nghia ra mot khung cua cac thuat toan trong lop
// co so va de cac lop con ghi de len cac buoc ma khong lam thay doi cau truc tong the cua project

//trong truong hop nay, chung ta co mot interface IOtp va mot struct Otp
// struct Otp co mot phuong thuc genAndSendOTP
// phuong thuc genAndSendOTP se goi cac phuong thuc khac cua struct Otp
// cac phuong thuc khac cua struct Otp se duoc ghi de boi cac struct khac
// cac struct khac se implement interface IOtp
// cac struct khac se co cac phuong thuc khac nhau
// cac phuong thuc khac nhau se duoc goi trong phuong thuc genAndSendOTP
// phuong thuc genAndSendOTP se duoc goi trong main

func main() {
	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &Email{}
	o = Otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)

}
