package main

import "fmt"

func main() {
	// deferm()
	x := defer2()
	fmt.Println(x)
	x = d4()
	fmt.Println(x)
}

func deferm() {
	fmt.Println("hello ")
	defer fmt.Println("world")
	defer fmt.Println("kdd")
	fmt.Println("welcome")
}

//go语言函数的return不是原子操作，底层是分为两步执行的
//1. 对返回值赋值
//2. 真正的return
//若函数中存在defer，那么defer执行的时机是在上述两步之间

//这个函数的结果会返回5
func defer2() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func defer3() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func d4() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func d5() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
