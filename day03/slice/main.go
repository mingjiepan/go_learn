package main

import "fmt"

//数组的局限性
//数组长度是固定的，并且长度是类型的一部分，那么当一个函数的参数为 [4]int时
//其他长度的数组值便无法传递
//还有数组长度是固定的，无法往数组添加更多元素，因此需要有切片
//切片是一个拥有相同类型的可变长度的序列，它是基于数组类型的一层封装，它非常灵活，支持自动扩容
//切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合
func main() {
	// slice1()
	// slice2()
	// slice3()
	// slice4()
	// slice5()
	// slice6()
	deleteSlice()
}

func slice1() {
	var s1 []int    //定义一个存放int类型元素的切片
	var s2 []string //定义一个存放String类型元素的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	fmt.Printf("s1 len %d s1 cap %d\n", len(s1), cap(s1))

	s1 = []int{1, 2, 3}
	fmt.Printf("s1 len %d s1 cap %d\n", len(s1), cap(s1))
}

//切片指向了一个底层的数组
//切片的长度就是元素的个数
//切片的容量是底层数组从切片的第一个元素到数组最后一个元素的数量
func slice2() {
	//由数组得到切片
	fmt.Println("------------------------------")
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4]
	s4 := a1[:4]
	s5 := a1[3:]
	s6 := a1[:]
	a1[3] = 300 //改变底层数组的值，同样也会反应到切片中
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)
	fmt.Printf("s3 len %d s3 cap %d\n", len(s3), cap(s3))
	fmt.Printf("s4 len %d s4 cap %d\n", len(s4), cap(s4))
	fmt.Printf("s5 len %d s5 cap %d\n", len(s5), cap(s5))
	fmt.Printf("s6 len %d s6 cap %d\n", len(s6), cap(s6))
}

//通过make函数，初始化切片
func slice3() {
	var s1 []int = make([]int, 5, 10)
	fmt.Printf("s1 %v len %d cap %d\n", s1, len(s1), cap(s1))
	var s2 []int = make([]int, 0, 10)
	fmt.Printf("s2 %v len %d cap %d\n", s2, len(s2), cap(s2))
}

//切片之间不能直接比较，不能使用 == 操作符 来比较
//切片唯一合法的比较操作是和nil比较，nil值的切片没有底层数组，一个nil值的切片长度和容量都是0
//但是我们不能说一个长度和容量都是0的切片一定是nil
//要判断一个切片是否是空的，用 len(s) == 0来判断，而不应该使用 s == nil 来判断
//切片是引用类型，指向底层的数组

func slice4() {
	s1 := make([]int, 5, 10)
	s1[7] = 10
	// s2 := make([]int, 5, 10)
	// fmt.Println(s1 == s2)
}

func slice5() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	// s1[3] = "广州"
	// fmt.Println(s1)
	s1 = append(s1, "广州")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}

//copy和赋值不同的，copy底层会新开辟新的空间，而赋值是共用的
func slice6() {
	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 3, 3)
	copy(s2, s1)
	s1[2] = 30
	fmt.Println(s2)
}

//切片没有删除元素的功能
func deleteSlice() {
	s1 := []int{1, 2, 3, 4}
	s1 = append(s1[:2], s1[3])
	fmt.Println(s1)

	//切片会修改底层数组
	s2 := [...]int{1, 2, 3, 4}
	s3 := s2[:]
	s3 = append(s3[:2], s3[3])
	fmt.Println(s2)
}
