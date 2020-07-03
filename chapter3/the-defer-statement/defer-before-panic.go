package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer!!!")
	}()

	// panic("paniced!!!")

	// 패닉 만들어봄
	a := []int{1}
	_ = a[3]
}

/*
defer 와 panic 어느게 먼저 발생할까?
defer 가 먼저 발생한다!

*/
