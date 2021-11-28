package main

import (
	"fmt"
	"sync"
)

//通道的三种操作
//发送
//c <- 10

//接收
//r := <- c
//忽略结果  <- c

//关闭
//close(c)
//关闭后不能写入，但还可以读取

//要注意区别有缓冲区的channel和没有缓冲区的channel的区别
var cc chan int
var wg1 sync.WaitGroup

func main() {
	wg1.Add(1)
	cc = make(chan int)

	go func() {
		defer wg1.Done()
		rr := <- cc
		fmt.Println(rr)
	}()

	//程序不会往下执行了
	cc <- 10

	fmt.Println("hello world")
	wg1.Wait()
	close(cc)
}