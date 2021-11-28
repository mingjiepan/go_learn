package main

import "fmt"

//值接受者和指针接受者的区别

type fruit interface {
	eat()
}

type apple struct {}

//指针接受者
func (a *apple) eat() {

}


func main() {
	var f fruit

	a := apple{}

	f = &a
	fmt.Println(f)
}
