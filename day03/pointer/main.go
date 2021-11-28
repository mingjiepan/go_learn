//go语言中不存在指针操作，只需要记住两个符号：
// &:取地址
// *:根据地址取值

//new 函数申请一个内存地址，一般是给基本数据类型申请内存
//make也是用于内存分配的，区别于new，只用于slice、map、chan的内存创建，且返回的是这三个类型
//本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以没必要返回他们的指针了
package main

import "fmt"

func main() {
	// p1()
	p2()
}

func p1() {
	a := 10
	b := &a
	fmt.Println(b)
	fmt.Println(&b)
	fmt.Printf("%T\n", b)
	fmt.Println(*b)
}

func p2() {
	// var a *int //nil
	// *a = 100   //执行时会报错
	// fmt.Println(*a)

	//new函数申请一个内存地址
	var b *int = new(int)
	*b = 200
	fmt.Println(b)
	fmt.Println(*b)
}
