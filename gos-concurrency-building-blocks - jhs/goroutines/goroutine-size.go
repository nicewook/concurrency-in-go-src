package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memStats := func() runtime.MemStats {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() {
		wg.Done()
		<-c
	}

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memStats()
	fmt.Printf("goroutine num: %v\n", runtime.NumGoroutine())
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memStats()
	fmt.Printf("%.3fkb\n", float64(after.Sys-before.Sys)/numGoroutines/1000)
	fmt.Printf("goroutine num: %v\n", runtime.NumGoroutine())
	// fmt.Printf("--\nbefore memStats\n%#v\n", before)
	// fmt.Printf("--\nafter memStats\n%#v\n", after)
}

/*
1) wg.Wait() 는 모든 고루틴 (1e4 개)가 다 돌때까지 기다리고
2) 고루틴은 그래도 죽지 않는다
3) 그러고 나서 10000 의 고루틴이 살아있을때의 sys 값 차이를 본다 Sys 는 OS 로부터 얻은 메모리 사이즈이다.
1e4 에서 8.836kb 결과는 겨우 8kb 밖에 안됨
1e5, 1e6 에서도 별 차이가 안남. 흠...
*/
