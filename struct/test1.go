package main

import "fmt"

//结构体
type person struct {
	name string
	age int
	gender int
	hobby []string
}

func main() {
	//声明一个person的类型
	var p person
	//通过字段赋值
	p.name = "kd"
	p.age = 10
	p.gender = 1
	p.hobby = []string{"篮球", "足球", "乒乓球"}
	fmt.Println(p)
	fmt.Println(p.name)

	//匿名结构体:临时使用的
	var s struct{
		name string
		age int
	}
	s.age = 10
	s.name = "kd"
	fmt.Println(s)
}