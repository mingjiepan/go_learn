package main

import "fmt"

//go语言中函数的参数永远都是拷贝
//结构体是值类型
//go中make和new都是用来申请内存地址，make用于运用类型的，new用于值类型的，new返回的是指针，make返回的是变量
//结构体内部的字段在内存上是一段连续的内存
type student struct {
	name string
	age int
	gender int
	hobby []string
}

//传的是拷贝
func f(x student) {
	//修改副本的值
	x.age = 20
}

//传的是指针
func f2(x *student) {
	(*x).age = 20 //根据内存地址找到那个原变量，修改的就是原来的变量
	x.name = "haha"//go的语法糖，会自动判断x是否为指针类型
}

func main() {
	var p student
	p.name =  "kobe"
	p.age = 10
	f(p)
	fmt.Println(p)
	f2(&p)
	fmt.Println(p)

	fmt.Println("---------1------------")

	var p2 = new(student)
	(*p2).gender = 2
	p2.name = "hello" //语法糖
	fmt.Printf("type p2 %T\n", p2)
	fmt.Printf("%p\n", p2)
	fmt.Println(p2)

	fmt.Println("--------2----------")

	//使用key，value初始化
	var s = student{
		name: "kg",
		gender: 2,
	}
	fmt.Println(s)
	fmt.Printf("s %p\n", &s)

	//使用值列表初始化
	s2 := student{
		"hh",
		2,
		2,
		[]string{"篮球", "足球"},
	}
	fmt.Println(s2)

	//初始化后，直接拿到指针
	s3 := &student{
		name: "gg",
		gender: 12,
	}
	fmt.Println(s3)
	fmt.Printf("s3 name %p s3 gender %p\n", &(s3.name), &(s3.gender))

	fmt.Println("-----------------------------")
}