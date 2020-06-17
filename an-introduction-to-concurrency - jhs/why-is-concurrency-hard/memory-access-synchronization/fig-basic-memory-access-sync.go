package main

import (
	"fmt"
	"sync"
)

func main() {
	var memoryAccess sync.Mutex
	var value int

	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if value == 0{
		fmt.Printf("value: %v\n", value)  // this will not print 1 never ever!
	} else {
		fmt.Printf("value: %v\n", value)
	}
	memoryAccess.Unlock()
}