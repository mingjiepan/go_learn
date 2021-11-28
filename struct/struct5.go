package main

import "fmt"

//结构体的"继承"
//通过匿名嵌套来实现继承
type animal struct {
	name string
	age int
}

func (a *animal) move() {
	fmt.Printf("%s moving...\n", a.name)
}

type cat struct {
	color string
	animal
}

func (c *cat) miao() {
	fmt.Printf("%s miao...\n", c.name)
}

func main() {
	c1 := cat{
		color: "red",
		animal : animal{
			name: "c1",
			age: 20,
		},
	}
	c1.move()
	c1.miao()
}