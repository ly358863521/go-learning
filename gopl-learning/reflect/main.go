package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	t := reflect.ValueOf(3)
	fmt.Println(t)          //3
	fmt.Println(t.Type())   //int
	fmt.Println(t.String()) //<int value>
	x := t.Interface()
	fmt.Println(x)       //3
	fmt.Println(x.(int)) //3

	p := reflect.ValueOf(make([]int, 3))
	fmt.Println(p)
	fmt.Println(p.Kind())
	fmt.Println(p.Type())
	fmt.Println(p.String())
	fmt.Println(p.Pointer())
	fmt.Println(uint64(p.Pointer()))
	fmt.Println("0x", strconv.FormatUint(uint64(p.Pointer()), 16))
	/*
		[0 0 0]
		slice
		[]int
		<[]int Value>
		824634122592
		0x 824634122592
		0x c000062160

	*/
}
