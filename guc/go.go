package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello world", i)
}

//程序创建之后会创建一个主goroutine来执行
func main() {
	for i := 0; i < 10000; i++ {
		go hello(i) //开启一个单独的goroutine来执行
	}
	fmt.Println("welcome")
	//main函数结束了，由main函数 启动的goroutine也都结束了
	time.Sleep(time.Second)
}

/**
通过匿名函数执行
 */
func world() {
	go func(){
		fmt.Println("welcome")
	}()
}