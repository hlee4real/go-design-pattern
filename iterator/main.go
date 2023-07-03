package main

import "fmt"

// iterator cho phep duyet tuan tu cac yeu cau thong qua mot cau truc du lieu phuc tap ma khong lam lo chi tiet ben trong no

//trong vi du nay, chung ta se tao ra mot iterator de duyet qua cac doi tuong user trong mot collection

func main() {
	user1 := &User{
		name: "a",
		age:  30,
	}
	user2 := &User{
		name: "b",
		age:  20,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
