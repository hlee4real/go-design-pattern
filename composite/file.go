package main

import "fmt"

// Composite là một mẫu thiết kế cấu trúc cho phép kết hợp các đối tượng thành một cấu
// trúc dạng cây và làm việc với cấu trúc đó như thể nó là một đối tượng đơn lẻ.
type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Println("Searching for keyword", keyword, "in file", f.name)
}

func (f *File) getName() string {
	return f.name
}
