package main

import "fmt"

func main() {
	//age 只能在ifelse作用域内访问，无法在之外访问
	// if age := 10; age > 10 {
	// 	fmt.Println("hello world")
	// } else {
	// 	fmt.Println("welcome")
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// j := 1
	// for ; j < 10; j++ {
	// 	fmt.Println(j)
	// }

	// for j < 10 {
	// 	fmt.Println(j)
	// 	j++
	// }

	//无限循环
	// for {

	// }

	str := "hello world 你好么"
	for index, value := range str {
		fmt.Println(index)
		fmt.Printf("%c\n", value)
	}
}
