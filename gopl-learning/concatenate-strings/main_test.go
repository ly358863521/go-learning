package main

import (

    "testing"
)
func Benchmark_f1(b *testing.B){
    b.N = 12345678
    for i := 0; i < b.N ; i++ {
		f1()
	}
}
func Benchmark_f2(b *testing.B){
    b.N = 12345678
    for i := 0; i < b.N ; i++ {
		f2()
	}
}
func Benchmark_f3(b *testing.B){
    b.N = 12345678
    for i := 0; i < b.N ; i++ {
		f3()
	}
}

func  Test_f1(t *testing.T){
    f1()
}

//go test -v ./main_test.go ./main.go -bench=".*"        
//只显示压力测试
//go test -v .