package main

import "fmt"

//匿名函数
var ff = func(x, y int) {
	fmt.Println(x, y)
}

func main() {
	a := f1
	b := f2
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	r := f3(f2)
	fmt.Println(r)
}

func f2() int {
	return 2
}

func f1() {
	fmt.Println("hello")
}

//函数也可以作为参数的类型
func f3(x func() int) int {
	return x()
}
