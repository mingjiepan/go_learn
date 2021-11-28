package main

import "fmt"

func main() {

	result := testSwitch1(2)
	fmt.Println(result)

	result = testSwitch2(5)
	fmt.Println(result)

	result = testSwitch3(12)
	fmt.Println(result)

	testGoto()
}

func testSwitch1(num int) string {
	switch num {
	case 1:
		return "hello"
	case 2:
		return "world"
	default:
		return "welcome"
	}
}

func testSwitch2(num int) string {
	switch num {
	case 1, 2, 3, 4:
		return "hello"
	case 5, 6:
		return "world"
	default:
		return "welcome"
	}
}

func testSwitch3(num int) string {
	switch {
	case num < 5:
		return "hello"
	case num > 10:
		return "world"
	default:
		return "welcome"
	}
}

func testGoto() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 9; j++ {
			fmt.Println("hello", j)
			if j == 6 && i == 2 {
				goto outer
			}
		}
	}
outer:
}
