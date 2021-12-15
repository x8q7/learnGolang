package main

import (
	"reflect"
)

func main() {
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	// typeOfT := s.Type() // 此处 转成 reflect.Type 类型
	for i := 0; i < s.NumField(); i++ {
		// f 是 reflect.Value 类型 的 字段的 value
		f := s.Field(i)

		if f.Kind() == reflect.Int {
			f.SetInt(10)
		}

		if f.Kind() == reflect.String {
			f.SetString("hi")
		}
	}
}
