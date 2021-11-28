package main

import (
	"fmt"
	"strings"
)

//go 语言中，字符串必须用双引号
// 单引号包的是字符

//单独的字母、汉字、符号表示一个字符

//go 语言的字符都是int32类型的，但是有两种别名
// byte，表示的是ASCII码的字符，一个字节表示的
// rune，表示的是中文、韩文等其他字符，一般是3个字节表示的

func main() {
	a := "你好么"
	b := '你'
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("a的长度%d \n", len(a))

	//多行字符串
	c := `
		hello world
		welcome
		hhahaha
	`
	fmt.Println(c)

	//字符串拼接
	province := "福建省"
	city := "厦门市"
	address := province + city
	fmt.Println(address)

	fmt.Printf("%s%s", province, city)

	result := fmt.Sprintf("%s%s", province, city)
	fmt.Println("result=", result)

	//分割
	strArr := strings.Split(result, "门")
	fmt.Println(strArr)

	//包含
	fmt.Println(strings.Contains(province, "福建"))

	//前缀
	fmt.Println(strings.HasPrefix(result, "福"))
	//后缀
	fmt.Println(strings.HasSuffix(result, "省"))

	//index join

	//
	test()
	editStr()
	char()
}

func test() {
	fmt.Println("--------------------")
	s := "hello world welcome 你好么"

	// for i := 0; i < len(s); i++ {
	// 	// fmt.Println(s[i])
	// 	fmt.Printf("%c\n", s[i])
	// }

	for _, c := range s {
		fmt.Printf("%c\n", c)
	}
}

//字符串是不能修改的
func editStr() {
	fmt.Println("--------------------")
	s1 := "白萝卜"
	s2 := []rune(s1) //把字符串强制转成一个rune切片
	s2[0] = '你'
	fmt.Println(string(s2))
}

func char() {
	fmt.Println("--------------------")
	c1 := "红"
	c2 := '红'
	c3 := '1'
	fmt.Printf("c1 = %T, c2 = %T, c3 = %T \n", c1, c2, c3)
}
