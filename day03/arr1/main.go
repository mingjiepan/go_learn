package main

import "fmt"

//数组长度不可以变化，存放的相同类型的元素
//数组的长度是数组类型的一部分
func main() {
	// arr1()
	arr4()
}

//数组的声明
func arr1() {
	//a1和a2是不同的类型，因此无法比较
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:%T a2:%T\n", a1, a2)
	fmt.Println(a1, a2)
}

//数组的初始化
//如果不初始化，默认元素都是零值（布尔值是flase,int是0，字符串是空字符串）
func arr2() {
	//方式1
	var a1 [3]bool = [3]bool{true, true, true}
	fmt.Println(a1)
	//方式2，根据初始值自动推断数组的长度是多少
	a2 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a2)
	//方式3，
	a3 := [5]int{1, 2}
	fmt.Println(a3)
	//方式4,根据索引来初始化
	a4 := [5]int{0: 1, 4: 3}
	fmt.Println(a4)
}

//数组遍历的两种方式
func arr3() {
	a := [...]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println("-------------------")

	for index, value := range a {
		fmt.Println("index=", index, "value=", value)
	}
}

//多维数组
//[[1 2] [3 4] [5 6]]
func arr4() {
	var a [3][2]int
	a = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
	}
	// fmt.Println(a)
	for _, v := range a {
		for _, v1 := range v {
			fmt.Println(v1)
		}
	}
}

//数组是值类型，赋值和参数会赋值整个数组
func arr5() {
	a1 := [3]int{1, 2, 3}
	a2 := a1
	a2[0] = 10
	fmt.Println(a1, a2)
}

func arr6() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	sum := 0
	for _, v := range a {
		sum = sum + v
	}
	fmt.Println(sum)
}
