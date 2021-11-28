package main

import "fmt"

//常量
const pi = 3.1415926

//批量声明常量
const (
	ok       = 200
	notFound = 404
)

//若某一行没有写值，那么默认和上一行是一样的
const (
	n1 = 100
	n2
	n3
)

//iota 在const关键字出现时将被重置为0，const中每新增一行常量声明，将常量值加1吗，可以理解为行索引
//在定义枚举值时比较好用
const (
	a1 = iota //0
	a2 = iota
	_
	a3
	a4
)

//多个常量在同一行
const (
	d1, d2 = iota + 1, iota + 2 // d1=1 d2 = 2
	d3, d4 = iota + 1, iota + 2 // d3=2 d4 = 3
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("-------------------")
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	fmt.Println("a4:", a4)
}
