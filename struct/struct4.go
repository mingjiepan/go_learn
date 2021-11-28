package main

import "fmt"

/**
	嵌套结构体
 */
type stu struct {
	name string
	age int
}

type tea struct {
	no int
	stuArr []stu
}

type par struct {
	name string
	stu//匿名嵌套，
}

func main() {
	stu1 := &stu{age: 10}
	stu2 := &stu{name: "kd"}

	strArr := []stu{*stu1, *stu2}
	t1 := &tea{no: 1, stuArr: strArr}
	fmt.Println(t1)
	fmt.Println("--------")

	p := &par{
		name: "kd",
		stu: stu{
			name: "aa",
			age: 20,
		},
	}

	fmt.Println(p.age)//使通过外部的变量，可以直接访问匿名嵌套结构体的变量,如 p.age的age并不是p的变量
}