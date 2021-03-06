package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// 序列化 结构体( 注意：  冒号前后 不能有空格)
	type Person struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		IdCard  string `json:"id_card"`
		Address string `json:"address"`
	}

	a := &Person{Name: "xxx", Age: 16, IdCard: "12345678909887676", Address: "beijing"}

	jby, err := json.Marshal(a)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jby))
	
	// 反序列化
	var newPerson Person

	jbyStr := string(jby)

	err = json.Unmarshal([]byte(jbyStr), &newPerson)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(newPerson)

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

