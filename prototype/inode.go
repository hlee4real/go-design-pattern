package main

//prototype la mot design pattern thuoc creational pattern,cho phep sao chep (clone) cac doi tuong, tham chi la sao chep
//cac doi tuong phuc tap ma khong can ghep noi voi cac lop cu the cua chung

//tat ca cac prototype classes phai co mot interface chung de sao chep ngay ca khi cac lop cu the
//cua chung khong duoc biet den. trong truong hop nay la Inode interface va clone method

type Inode interface {
	print(string)
	clone() Inode
}
