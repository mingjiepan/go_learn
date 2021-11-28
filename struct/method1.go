package main
import (
	"fmt"
)

//方法：作用于特定类型的函数，记住与函数的区别
type car struct {
	price int
	name string
}

func newCar(price int, name string) *car {
	return &car{
		price: price,
		name: name,
	}
}

//值接受者和指针接受者
//接受者表示的是调用该方法的具体类型变量，多用类型首字母小写表示
func (d car)run() {
	fmt.Printf("car %s run.....\n", d.name)
}

//用值接受者，方法内部该结构体的字段，外部是看不到的
func (c car) hello() {
	c.price += 10
}

func (c *car) world() {
	(*c).price += 10
}

func main() {
	c1 := newCar(10, "宝马")
	c1.run()
	c1.hello()
	fmt.Println(c1.price)
	c1.world()
	fmt.Println(c1.price)
}