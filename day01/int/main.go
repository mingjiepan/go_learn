package main

import "fmt"

// %s 表示字符串
// %d 表示十进制
// %b 表示二进制
// %o 表示八进制
// %x 表示十六进制

// %f 表示float类型
// %.3f 表示float类型，保留小数点3位

// %T 查看变量的类型

// %v 打印变量的值，无论变量是什么值的
// %#v 打印变量的值，无论变量是什么值的，打印出来的还会带符号

func main() {
	a := 10
	fmt.Printf("a = %d \n", a)
	fmt.Printf("a = %b \n", a)
	fmt.Printf("a = %o \n", a)
	fmt.Printf("a = %x \n", a)

	//八进制,以0开头
	var b int = 077
	fmt.Printf("b = %o \n", b)

	//16进制，以0x开头
	var c int = 0xff
	fmt.Printf("c = %x \n", c)

	//查看变量的类型
	fmt.Printf("c的类型=%T \n", c)

	f1 := 3.12134
	fmt.Printf("f1 = %f \n", f1)
	fmt.Printf("f1 = %.2f \n", f1)
	fmt.Printf("f1 = %#v \n", f1)
}

func typeConvert() {
	var a int64 = 10
	var b int32 = int32(a)
	fmt.Println(b)
}
