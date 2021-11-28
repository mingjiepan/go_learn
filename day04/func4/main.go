package main

import "fmt"

//闭包是一个函数，这个函数包含了他外部作用域的一个变量
//底层的原理
//1.函数可以作为返回值
//2.函数内部查找变量的顺序，先在自己内部找，找不到往外层找

func f1(f func()) {
	f()
}

func f2(x, y int) {
	fmt.Println(x + y)
}

//需求，要将函数f2作为参数传递给f1,
// f2的类型和f1的不一样，因此不能直接传递，需要通过第三个函数进行封装处理

//定义一个函数对f2进行包装
func f3(x int, y int, f func(int, int)) func() {
	return func() {
		f(x, y)
	}
}

func main() {
	f1(f3(1, 2, f2))
}

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
