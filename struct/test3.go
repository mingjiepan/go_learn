package main

import "fmt"

type x struct {
	a int8
	b int8
	c int8
}

//结构体内部变量占用的内存是连续的，不间断的

func main() {
	a := &x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p\n", &(a.a))
	fmt.Printf("%p\n", &(a.b))
	fmt.Printf("%p\n", &(a.c))
}
