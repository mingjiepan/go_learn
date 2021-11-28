package main

import (
	"fmt"
	"os"
)

func writeFile() {
	file, err := os.OpenFile("C:\\Users\\kd\\Desktop\\hello.txt", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file error %v", err)
		return
	}
	defer file.Close()
	file.WriteString("hello world")
}

func main() {
	writeFile()
}

