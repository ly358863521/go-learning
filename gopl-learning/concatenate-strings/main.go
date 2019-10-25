package main

import (
	// "fmt"
	"strconv"
	"strings"
	"bytes"
)
func f1(){
	s := ""
    for i := 0; i <= 9; i++ {
		s += strconv.Itoa(i)
		//s += string(i)
    }
	//fmt.Println(s)
}
func f2(){
	var s []string
    for i := 0; i <= 9; i++ {
        s = append(s, strconv.Itoa(i))
    }
	// fmt.Println(strings.Join(s, ""))
	strings.Join(s,"")
}
func f3(){
	var buffer bytes.Buffer
    for i := 0; i <= 9; i++ {
        buffer.WriteString(strconv.Itoa(i))
    }
	// fmt.Println(buffer.String())
	buffer.String()
}
func main(){
	f1()
	f2()
	f3()
}