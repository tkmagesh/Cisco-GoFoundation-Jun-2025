package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for range 20 {
		wg.Add(1)
		go f1(wg)
	}
	f2()
	wg.Wait()
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("f1 started")
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Println("f1 completed")

}

func f2() {
	fmt.Println("f2 invoked")
}
