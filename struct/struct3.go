package main

import "fmt"

//匿名字段，实际开发不建议使用
type mm struct {
	int
	string
	age int
}



func main() {
	m := &mm{
		10,
		"zhangsan",
		30,
	}
	fmt.Println(m)
	fmt.Println(m.int)
	fmt.Println(m.string)
	fmt.Println(m.age)
}
