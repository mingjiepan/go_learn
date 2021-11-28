package main

import "fmt"

func main() {
	m2()
}

func m1() {
	var m1 map[string]string = make(map[string]string, 10) // 估算使用空间，避免扩容
	m1["address"] = "厦门"
	fmt.Printf("m1= %v len(m1)= %d\n", m1, len(m1))

	fmt.Println(m1["age"]) //如果map不存在这个key，那么拿到的是value类型对应的零值
	value, ok := m1["name"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("没查到此键")
	}

	fmt.Println("------------------------------------")

	for k, v := range m1 {
		fmt.Printf("k=%v v=%v\n", k, v)
	}

	fmt.Println("------------------------------------")

	//删除map的key
	delete(m1, "name")

	fmt.Println("------------------------------------")
}

//map 和 slice的组合使用
func m2() {
	// var s1 = make([]map[int]string, 0, 10)
	// s1[0][1] = "xiamen"
	// fmt.Println(s1)
	var m1 = make([]map[int]string, 10, 10)
	m1[0] = make(map[int]string, 1)
	m1[0][10] = "xiamen"
	fmt.Println(m1)
}
