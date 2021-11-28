package main

import (
	"fmt"
)

//接口是一个类型
//不关心变量是什么类型，只关心变量能调用什么方法
type speaker interface {
	speak()
}

type person struct {}

type cat struct {}

type dog struct {}

func (c cat) speak() {
	fmt.Println("喵喵喵...")
}

func (d dog) speak() {
	fmt.Println("汪汪汪...")
}

func (p person) speak() {
	fmt.Println("hello world")
}

func speak(s speaker) {
	s.speak()
}

func move(s speaker) {

}

func main() {
	var s speaker
	p := person{}
	s = p

	p2 := &person{}
	s = p2

	speak(s)
	speak(p2)
}
