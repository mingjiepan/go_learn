package main

import "fmt"

type ss struct {
	name string
	age int
}

//手动写一个构造函数，一般使用new开头
//什么时候返回结构体类型，什么时候返回结构体指针
//当结构体比较小，拷贝的内存开销不大，可以返回结构体类型，当结构体类型大的时候，拷贝的内存开销大，适合返回指针
//建议返回指针
func newSs(name string, age int) *ss {
	return &ss{
		name: name,
		age: age,
	}
}

func main() {
	s1 := newSs("kd", 1)
	s2 := newSs("dd", 2)
	fmt.Println(s1, s2)
}