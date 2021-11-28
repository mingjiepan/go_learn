package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

//goroutine什么时候结束？对应的函数结束了，goroutine就结束了
//main函数执行完了，由main函数的那些goroutine也就跟着结束了
func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go f1(i)
	}
	wg.Wait()
	fmt.Println("所有的goroutine都执行完了")
}

func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r := rand.Intn(10)
		fmt.Println(r)
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	fmt.Println("hello", i)
}