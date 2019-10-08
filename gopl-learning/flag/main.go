package main

import (
    "flag"
    "fmt"
)

//定义一个字符串变量，并制定默认值以及使用方式
var species = flag.String("species", "gopher", "the species we are studying")

//定义一个int型字符
var nums  = flag.Int("ins", 1, "ins nums")

func main() {
  // 上面定义了两个简单的参数，在所有参数定义生效前，需要使用flag.Parse()来解析参数
 flag.Parse()
  // 测试上面定义的函数
 fmt.Println("a string flag:",string(*species))
 fmt.Println("ins num:",rune(*nums))
}

// $ go run flag-test.go 
// a string flag: gopher
// ins num: 1

// $ go run flag-test.go -ins 3 -species biaoge
// a string flag: biaoge
// ins num: 3
