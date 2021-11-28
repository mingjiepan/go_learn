package main

import "fmt"

//go通过通信共享内存，而不是通过共享内存而实现通信
//go的channel是一种特殊的类型，goroutine之间的连接，像一个队列，总是遵循先入先出的规则，保证收发数据的顺序
//声明channel的时候需要为其指定channel的元素类型

var c chan int //指定通道中的元素类型
var d chan []int

func main() {
	c = make(chan int) //通道的初始化
	d = make(chan []int, 10)//带缓冲区大小的初始化
	fmt.Println(c)
	fmt.Println(d)


}