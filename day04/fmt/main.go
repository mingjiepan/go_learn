package main

import "fmt"

/*

fmt.Prinf
%T 查看类型
%d 十进制
%o 八进制
%b 二进制
%x 十六进制
%c zifu
%s 字符串
%p 指针
%v 值
%f 浮点数
%t 布尔值

通用占位符
%T
%+v
%#v
%v
%%
*/

func main() {
	// var s string
	// fmt.Scan(&s)
	// fmt.Println("用户输入的内容是", s)

	var (
		name  string
		age   int
		class string
	)
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Println(name, age, class)
}

func pf() {
	m := 123.456789
	fmt.Printf("%.2f\n", m)
}
