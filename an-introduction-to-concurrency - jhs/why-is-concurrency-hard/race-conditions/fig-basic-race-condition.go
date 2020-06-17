// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var data int
// 	go func() { // <1>
// 		data++
// 	}()
// 	if data == 0 {
// 		fmt.Printf("the value is %v.\n", data)
// 	}
// }

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
	f, err := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	for i := 0; i < repeat; i++ {
		raceCond(f)
	}
	defer f.Close()
}

// possible
// data: 0, the value is 0.
// data: 0, the value is 1.
// data: 1, the value is 1.
