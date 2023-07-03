package main

import "fmt"

//prototype la mot design pattern thuoc creational pattern,cho phep sao chep (clone) cac doi tuong, tham chi la sao chep
//cac doi tuong phuc tap ma khong can ghep noi voi cac lop cu the cua chung

//tat ca cac prototype classes phai co mot interface chung de sao chep ngay ca khi cac lop cu the
//cua chung khong duoc biet den. trong truong hop nay la Inode interface va clone method

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}
