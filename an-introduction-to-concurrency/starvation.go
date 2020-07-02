/*
greedyWorker 와 politeWorker 두 고루틴이
누가 더 많이 도는지를 비교하는 코드였다.

1. 재미있었던 것은 Windows 에서는 대략 1ms 이하로는 sleep 에 들어가지 않는다는 것이었다.
	- 그보다 짧게 Sleep 을 주면 대부분 비슷한 결과가 나오게 된다.
	- 따라서 코드상으로 각각 3ns 과 1ns *3 을 자게 하면 greedy 는 딱 한 번 1ms 을 잔다면, polite 는 1ms 을 세 번이나 자게 된다.
		Greedy worker was able to execute 341 work loops
		Polite worker was able to execute 114 work loops.
	- 하지만 코드상으로 각각 6ms 과 1ms *3 을 자게 하면 greedy 는 딱 한 번 1ms 을 잔다면? - 비슷하넹. 하하
	  Greedy worker was able to execute 113 work loops
	  Polite worker was able to execute 38 work loops.

2. 아무튼 한 번의 루프에 똑같은 sleep 시간을 주어도 Lock 을 덜 빼앗기고 더 빼앗을 수 있는 greedy 가 더 많이 루프를 돈다는 것이다.



*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second

	greedyWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			// time.Sleep(3 * time.Nanosecond)
			time.Sleep(9 * time.Millisecond)
			sharedLock.Unlock()
			count++
		}

		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}

	politeWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			// time.Sleep(1 * time.Nanosecond)
			time.Sleep(3 * time.Millisecond)
			sharedLock.Unlock()
			sharedLock.Lock()
			// time.Sleep(1 * time.Nanosecond)
			time.Sleep(3 * time.Millisecond)
			sharedLock.Unlock()
			sharedLock.Lock()
			// time.Sleep(1 * time.Nanosecond)
			time.Sleep(3 * time.Millisecond)
			sharedLock.Unlock()
			count++
		}

		fmt.Printf("Polite worker was able to execute %v work loops. \n", count)
	}

	wg.Add(2)
	go greedyWorker()
	go politeWorker()

	wg.Wait()
}
