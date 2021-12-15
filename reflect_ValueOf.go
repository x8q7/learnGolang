package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x) // Note: take the address of x.
	p = p.Elem()

	p.SetFloat(1.2)

	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	fmt.Println("reInterface", p.Interface())
}
