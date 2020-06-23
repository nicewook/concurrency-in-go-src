package main

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}

	sender := func() {
		defer wg.Done()
		<-begin // block
		for i := 0; i < b.N; i++ {
			c <- token // send via channel
		}
	}

	receiver := func() {
		defer wg.Done()
		<-begin // block
		for i := 0; i < b.N; i++ {
			<-c // send via channel
		}
	}

	wg.Add(2)
	go receiver()
	go sender()
	b.StartTimer() // manually start benchmark timer here
	close(begin)   // trigger
	wg.Wait()
}

/*
go test context_test.go -bench .
goos: windows
goarch: amd64
BenchmarkContextSwitch-12        5999942               204 ns/op
PASS
ok      command-line-arguments  1.591s

오히려 여러개의 CPU 를 돌아다니며 context switching 하는것 보다 이게 더 빠르네 하하
go test context_test.go -bench . -cpu=1
goos: windows
goarch: amd64
BenchmarkContextSwitch   9370021               126 ns/op
PASS
ok      command-line-arguments  1.464s
*/
