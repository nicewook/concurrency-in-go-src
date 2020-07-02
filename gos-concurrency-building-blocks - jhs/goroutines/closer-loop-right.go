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

// go func 에 그때 그때 i 값을 파라미터로 넘겨주면 문제 해결
