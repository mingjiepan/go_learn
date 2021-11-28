package main

import (
	"fmt"
	"sync"
)

//单向通道

// chan<- int
// <-chan int

var wg4 sync.WaitGroup
var wg5 sync.WaitGroup

func mm1(ch chan<- int) {
	defer wg4.Done()
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func mm2(ch <-chan int) {
	wg4.Wait()
	defer wg5.Done()
	for {
		r,ok := <- ch
		if !ok {
			break
		}
		fmt.Println(r)
	}
}

func main() {
	wg4.Add(1)
	wg5.Add(1)

	ch := make(chan int, 100)

	go mm1(ch)
	go mm2(ch)

	wg5.Wait()
	fmt.Println("结束了")
}