/*
sync.Mutex Lock() 안의 변수(==메모리) 에는
Unlock() 되기전에는 접근할 수 없는 것이다.

따라서 go func() 가 실행되면 value: 1 이 찍히겠지만
보통은 그전에 이미 main 고루틴이 실행되고 끝나버릴 것이다.

이것도 10,000번 반복하게 해보면 go func() 가 먼저 실행되는 경우가 나온다

*/
package main

import (
	"fmt"
	"os"
	"sync"
)

func mutexTest(f *os.File) {
	var memoryAccess sync.Mutex
	var value int
	var str string

	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if value == 0 {
		str = fmt.Sprintf("mutexTest - value: %v\n", value) // this will not print 1 never ever!
	} else {
		str = fmt.Sprintf("mutexTest - value: %v\n", value)
	}
	memoryAccess.Unlock()

	if _, err := f.WriteString(str); err != nil {
		panic(err)
	}
}

func main() {

	repeat := 10000
	f, err := os.OpenFile("memory-access-sync.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	for i := 0; i < repeat; i++ {
		mutexTest(f)
	}
	defer f.Close()

}
