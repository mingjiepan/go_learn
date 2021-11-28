package main

import "fmt"

//panic 和 recover
//recover必须搭配defer使用，用于恢复现场
//defer一定要在可能引发panic的语句之前定义

func m1() {
	fmt.Println("a")
}

func m2() {
	defer func() {
		fmt.Println("释放数据库连接")
		err := recover()
		if err != nil {
			//真正释放连接
		}
	}()
	//刚刚打开一个数据库连接
	panic("出错啦")
	fmt.Println("b")
}

func m3() {
	fmt.Println("c")
}

func main() {
	m1()
	m2()
	m3()
}
