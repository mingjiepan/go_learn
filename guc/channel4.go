package main

import (
	"fmt"
	"sync"
)

// channel练习

//1. 启动一个goroutine，生成100个数发送到chan1
//2. 启动一个goroutine，从chan1中取值，计算其平方发送到ch2中


var once sync.Once
var wg3 sync.WaitGroup

func m1(ch chan int) {
	defer wg3.Done()
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func m2(ch1 , ch2 chan int) {
	defer wg3.Done()
	for {
		x,ok := <- ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}

	//确保多个goroutine来执行时，只会执行一次
	once.Do(func() {
		close(ch2)
	})
}

func main() {

	wg3.Add(3)


	a := make(chan int, 100)
	b := make(chan int, 100)


	//一个生产者，两个消费者
	go m1(a)
	go m2(a, b)
	go m2(a, b)

	wg3.Wait()

	for r := range b {
		fmt.Println(r)
	}
}
