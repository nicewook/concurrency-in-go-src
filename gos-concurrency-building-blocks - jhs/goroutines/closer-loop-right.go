package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	max := 10000
	wg.Add(max)
	for i := max; i > 0; i-- {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

// go func 가 실행되는 타이밍에서의 i 값이 먹히는 거다.
