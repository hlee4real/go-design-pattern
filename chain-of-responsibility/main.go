package main

// chain of responsibility la 1 behavioral design pattern
// no cho phep cac doi tuong co kha nang xu ly cac yeu cau duoc truyen qua cho cac doi tuong lien quan cho toi khi
// mot yeu cau duoc xu ly thanh cong

//trong truong hop nay chung ta se xay dung mot he thong benh vien
// cac phong ban trong benh vien se duoc ket noi voi nhau theo mot thu tu nhat dinh
// khi mot benh nhan den benh vien, benh nhan se duoc gui toi phong tiep tan
// phong tiep tan se gui benh nhan toi phong kham
// phong kham se gui benh nhan toi phong kham benh
// phong kham benh se gui benh nhan toi phong thu ngan
// phong thu ngan se thu tien va ket thuc qua trinh

func main() {

	cashier := &Cashier{}

	//Set next for medical department
	medical := &Medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}
