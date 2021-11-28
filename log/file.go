package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//readByBufio()
	readByIoutil()
}

func readBySlice() {
	file, err := os.Open("F:\\go_path\\src\\github.com\\mingjiepan\\helloworld\\log\\file.go")
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer file.Close()
	//读文件
	var tmp = make([]byte, 128)
	for {
		readCnt, err := file.Read(tmp)
		if err != nil {
			break
		}
		//fmt.Println("此次读取字节数 ", readCnt)
		fmt.Println(string(tmp[:readCnt]))
		if readCnt == 0 {
			break
		}
	}
}

//利用bufio这个包读取文件
func readByBufio() {
	file, err := os.Open("F:\\go_path\\src\\github.com\\mingjiepan\\helloworld\\log\\file.go")
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer file.Close()
	//创建读对象
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line error: %v", err)
		}
		fmt.Print(line)
	}
}

func readByIoutil() {
	file, err := os.Open("F:\\go_path\\src\\github.com\\mingjiepan\\helloworld\\log\\file.go")
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("read data error : %v", err)
		return
	}
	fmt.Println(string(bytes))
}