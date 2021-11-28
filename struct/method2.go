package main

import "fmt"

//给自定义类型添加方法
//不能给别的包里面的类型添加方法，只能给自己包里的类型添加方法
//要给别的包的类型添加方法，只能通过定义一个新类型如下面这个
type myInt int

func (m myInt) hello() int {
	fmt.Println("我是一个int啊。。。。")
	return int(m + 10)
}

func main() {
	var a int32
	a = 10
	var b int32 = 20
	c := int32(30)
	fmt.Println(a, b, c)

	fmt.Println("-----------------------")
	var m myInt = 10
	r := m.hello()
	fmt.Println(r)

	var n = myInt(20)
	fmt.Println(n)
}