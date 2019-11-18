package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	buf := []byte{67, 66, 65, 68, 69, 70}
	buf2 := []byte{0, 67, 67, 67, 67, 67}
	fmt.Println(&buf)
	fmt.Println(unsafe.Pointer(&buf))
	fmt.Println(unsafe.Pointer(&*(*string)(unsafe.Pointer(&buf))))

	fmt.Println(*(*string)(unsafe.Pointer(&buf2)))
	fmt.Println(reflect.ValueOf(*(*string)(unsafe.Pointer(&buf2))).Kind())
	fmt.Println(*(*string)(
		unsafe.Pointer(
			&reflect.StringHeader{
				Data: (*reflect.SliceHeader)(unsafe.Pointer(&buf2)).Data + uintptr(1),
				Len:  5})))

	fmt.Println(reflect.ValueOf(buf).Kind())
	fmt.Println(buf)

	s := "abcdefg"
	str1 := s[:4]
	str2 := s[:5]
	header1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	header2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))
	fmt.Println(header1.Data, header2.Data, reflect.ValueOf(str1).Kind())

	// &[67 66 65 68 69 70]
	// 0xc000004480
	// 0xc000004480
	//  CCCCC
	// string
	// CCCCC
	// slice
	// [67 66 65 68 69 70]
}
