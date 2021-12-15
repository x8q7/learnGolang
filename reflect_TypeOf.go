package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int
}

func main() {

	a := Person{
		Name: "bob",
		Age:  18,
	}

	rtyp := reflect.TypeOf(a)

	// fmt.Println(rtyp.NumField())

	field1, _ := rtyp.FieldByName("Name")

	/*
			Name string          // 字段名
		    PkgPath string       // 字段路径
		    Type      Type       // 字段反射类型对象
		    Tag       StructTag  // 字段的结构体标签
		    Offset    uintptr    // 字段在结构体中的相对偏移
		    Index     []int      // Type.FieldByIndex中的返回的索引值
		    Anonymous bool       // 是否为匿名字段
	*/

	fmt.Println(field1.Name)
	fmt.Println(field1.PkgPath)
	fmt.Println(field1.Type)
	fmt.Println(field1.Tag)
	fmt.Println(field1.Offset)
	fmt.Println(field1.Index)
	fmt.Println(field1.Anonymous)
}
