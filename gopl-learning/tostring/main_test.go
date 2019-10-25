package main

import (
	"time"

	"testing"
)

func BenchmarkConvertToTitle(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		_ = ConvertToTitle(i^16777217)
	}
}
func BenchmarkConvertToTitleStringConcat(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		_ = ConvertToTitleStringConcat(16777217)
	}
}

func BenchmarkConvertToTitlePubBuf(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		_ = ConvertToTitlePubBuf(16777217)
	}
}

func TestCalc(t *testing.T) {
	time.Sleep(time.Second * 2)
	t.Log(ConvertToTitle(166))
}
//go test -v ./main_test.go ./main.go -bench=".*"        
//只显示压力测试

