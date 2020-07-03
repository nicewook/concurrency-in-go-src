package main

import (
	"fmt"
	"sync"
)

func main() {
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var goRoutineRunning sync.WaitGroup
		goRoutineRunning.Add(1)
		go func() {
			goRoutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn() // 출력
		}()
		goRoutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Display dialog box")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast()
	clickRegistered.Wait()
}

// Mouse clicked
// Maximizing window
// Display dialog box

/*
1) subscribe 를 해주면 sync.Cond 를 통신수단으로 전달해주고, 실행할 함수 하나 - 여기서는 fmt 출력하나 하는 함수 전달해준다
2) 3 개의 subscribe 를 해주면 결국 각각 고루틴 하나씩, 총 3개가 실행되는데 이건 c.Wait() 에서 대기를 탄다
3) 마침내 button.Clicked 가 Broadcast 해주면 그 결과로 모든 c.Wait() 가 풀리면서 출력이 일어난다. 어느게 먼저 실행될지는 모른다.



*/
