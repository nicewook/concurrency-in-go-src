/*
greedyWorker 와 politeWorker 두 고루틴이
누가 더 많이 도는지를 비교하는 코드였다.

재미있었던 것은 Windows 에서는 대략 1ms 이하로는 sleep 에 들어가지 않는다는 것이었다.
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
