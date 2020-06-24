package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	rmFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Remove from Queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		// for len(queue) == 2 {
		// 	c.Wait()
		// }
		if len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, 1) // interface{} 타입의 슬라이스 이므로 뭐든 넣어도 상관없음
		go rmFromQueue(1 * time.Second)
		c.L.Unlock()
	}
	fmt.Printf("queue len: %v\n", len(queue)) // 2가 된다.

}

/*
1) queue 의 길이가 2가 될때까지 add 된다.
2) 2가 되면 Wait 가 걸린다. 그 타이밍에 queue 가 remove 하나 되고
3) 다시 add 된다.
* 음... 이건 결국 모든걸 다 지우진 못하네
*/
