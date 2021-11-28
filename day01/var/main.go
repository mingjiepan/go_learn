package main

import "fmt"

// var name string
// var age int
// var isOk bool

//声明一个变量同时赋值
var city string = "厦门"

//类型推导
var province = "福建"

//批量声明
var (
	name string
	age  int
	isOk bool
)

//go语言中非全局变量声明后，必须使用，不然编译不过去
func main() {

	//简短变量声明，只能用于函数内，不可声明为全局变量
	address := "思明区"

	name = "hello world"
	age = 11
	isOk = true

	fmt.Printf("name:%s \n", name)
	fmt.Println(age)
	fmt.Println(isOk)
	fmt.Println(address)

	//匿名变量，不会分配内存
	_, name = userInfo()
	fmt.Println(name)
}

func userInfo() (int, string) {
	return 10, "zhangsan"

}
