package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	salut := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(salut)

	}()
	wg.Wait()
	fmt.Println(salut)

}
