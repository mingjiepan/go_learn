package main

import (
	"encoding/json"
	"fmt"
)

type jj struct {
	Name string `json:"name" db:"name"`
	Age int 	`json:"age"`
}

func main() {
	j1 := jj{
		"zhangsan",
		10,
	}
	b,error := json.Marshal(j1)
	if error != nil {
		fmt.Println("json.Marshal error")
		return
	}
	fmt.Println(string(b))

	var j2 jj
	str := `{"name":"zhangsan","age":10}`
	error = json.Unmarshal([]byte(str), &j2)
	if error != nil {
		fmt.Println("Unmarshal error")
		return
	}
	fmt.Println(j2)
}