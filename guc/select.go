package main

import "fmt"

func select1(ch1, ch2, ch3 <-chan int) {
	select {
	case <- ch1:
		fmt.Println("channel1 data")
	case <- ch2:
		fmt.Println("channel2 data")
	case r := <-ch3:
		fmt.Println(r)
	default:
		fmt.Println("no receive data from channel")
	}
}

func main() {

}
