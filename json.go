package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// 序列化 结构体
	type Person struct {
		Name    string
		Age     int
		IdCard  string
		Address string
	}

	a := &Person{Name: "xxx", Age: 16, IdCard: "12345678909887676", Address: "beijing"}

	jby, err := json.Marshal(a)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jby))

	// 序列化 map

	var amp map[string]interface{} = make(map[string]interface{})

	amp["name"] = "xxx"
	amp["age"] = 12
	amp["address"] = "beijing"

	mpstr, err := json.Marshal(amp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(mpstr))

	// 序列化 切片

	var arr []interface{}

	arr = append(arr, map[string]interface{}{"name": "xxx", "age": 12, "address": "asdfasdf"})
	arr = append(arr, map[string]interface{}{"name": "qqq", "age": 11, "address": "qqqqqqq"})

	arrStr, err := json.Marshal(arr)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arrStr))

}
