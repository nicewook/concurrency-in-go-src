/*
레이스컨디션이 일어나게 만들어보자

가능한 결과
data: 0, the value is 0.
data: 0, the value is 1.
data: 1, the value is 1.

raceCond() 라는 함수를 repeat 만큼 실행시키는 코드이다.
결과는 race-condition.log 파일에서 확인 가능하다.
*/

package main

import (
	"fmt"
	"os"
)

func raceCond(f *os.File) {
	var data int
	go func() {
		data++
	}()

	var str string
	if data == 0 {
		str = fmt.Sprintf("data: 0, the value is %v.\n", data)
	} else {
		str = fmt.Sprintf("data: 1, the value is %v.\n", data)
	}
	if _, err := f.WriteString(str); err != nil {
		panic(err)
	}
}
func main() {
	repeat := 10000
	f, err := os.OpenFile("race-condition.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	for i := 0; i < repeat; i++ {
		raceCond(f)
	}
	defer f.Close()
}
