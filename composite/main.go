package main

// Composite là một mẫu thiết kế cấu trúc cho phép kết hợp các đối tượng thành một cấu
// trúc dạng cây và làm việc với cấu trúc đó như thể nó là một đối tượng đơn lẻ.
func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &Folder{
		name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("hoang")
}
