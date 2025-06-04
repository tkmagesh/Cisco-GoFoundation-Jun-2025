package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan []int)
	go generateNos(ch)
	nos := <-ch
	for _, no := range nos {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

func generateNos(ch chan []int) {
	nos := make([]int, 0)
	for i := range 10 {
		time.Sleep(500 * time.Millisecond)
		nos = append(nos, (i+1)*10)
	}
	ch <- nos
}
