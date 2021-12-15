package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 声明一个空结构体
	type cat struct {
		Name string
		// 带有结构体tag的字段
		Type int `json:"type" id:"100"`
	}
	// 创建cat的实例
	ins := cat{Name: "mimi", Type: 1}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	// 遍历结构体所有成员
	for i := 0; i < typeOfCat.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfCat.Field(i)
		// 输出成员名和tag
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}
	// 通过字段名, 找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		// 从tag中取出需要的tag
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}

	/*
			Name string          // 字段名
		    PkgPath string       // 字段路径
		    Type      Type       // 字段反射类型对象
		    Tag       StructTag  // 字段的结构体标签
		    Offset    uintptr    // 字段在结构体中的相对偏移
		    Index     []int      // Type.FieldByIndex中的返回的索引值
		    Anonymous bool       // 是否为匿名字段
	*/
}

