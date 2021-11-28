package main

import "fmt"

//空接口没必要起名字，通常定义如下
//interface{}
//两个场景使用 1， 函数的参数 2 map的值

//类型断言  x.()


func getType(t interface{}) {
	r,ok := t.(int)
	if !ok {
		fmt.Println("错了..")
	} else {
		fmt.Println("真实值", r)
	}
}

func getType2(a interface{}) {
	switch t := a.(type) {
	case string:
		fmt.Println("字符串", t)
	case int:
		fmt.Println("整形", t)
	}
}

func main() {
	var a interface{}
	a = 1
	a = "string"
	a = true
	a = []int{1,2,3}
	fmt.Println(a)
}