package main

// observer cho phep cac doi tuong theo doi su thay doi cua cac doi tuong khac

// trong truong hop nay la cac khach hang se theo doi su thay doi cua cac san pham
// khi san pham co su thay doi thi se gui email cho cac khach hang bang cach thong bao cho cac khach hang (subcribe)
// cac item available in stock cung se dc thong bao cho cac khach hang
// cac doi tuong o day la khach hang, item, interface la observer
func main() {

	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
