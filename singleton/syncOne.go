package main

import (
	"fmt"
	"sync"
)

/*
Singleton là một mẫu thiết kế creational, đảm bảo rằng chỉ có một đối tượng thuộc loại này tồn tại
và cung cấp một điểm truy cập duy nhất vào đối tượng đó cho bất kỳ code nào khác.

Singleton có những ưu và nhược điểm gần giống như các biến toàn cục.
Mặc dù chúng rất tiện dụng, nhưng chúng phá vỡ tính module của code của bạn.

Bạn không thể chỉ sử dụng một lớp phụ thuộc vào Singleton trong một số context
khác mà không chuyển Singleton sang context khác.
*/

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
