package main

import "fmt"

//go函数没有默认参数
func main() {
	m, n := m3()
	fmt.Println(m, n)
}

func m2() int {
	ret := 3
	return ret
}

func m(a int, b int) (ret int) {
	ret = a + b
	return //使用命名返回值 return后可以省略
}

func m3() (int, int) {
	return 1, 2
}

func m4(a, b int) {
	fmt.Println(a, b)
}

func m5(x string, y ...int) {

}
