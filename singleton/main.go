package main

import "fmt"

/*
Singleton là một mẫu thiết kế creational, đảm bảo rằng chỉ có một đối tượng thuộc loại này tồn tại
và cung cấp một điểm truy cập duy nhất vào đối tượng đó cho bất kỳ code nào khác.

Singleton có những ưu và nhược điểm gần giống như các biến toàn cục.
Mặc dù chúng rất tiện dụng, nhưng chúng phá vỡ tính module của code của bạn.

Bạn không thể chỉ sử dụng một lớp phụ thuộc vào Singleton trong một số context
khác mà không chuyển Singleton sang context khác.
*/

func main() {

	for i := 0; i < 30; i++ {
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
