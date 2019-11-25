package main

import (
	"fmt"
	"sort"
)

type Freq struct {
	key   int
	value int
}
type FreqList []Freq

func (p FreqList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p FreqList) Len() int           { return len(p) }
func (p FreqList) Less(i, j int) bool { return p[i].value < p[j].value } //递增排序
// func (p FreqList) Less(i, j int) bool { return p[i].value > p[j].value } 递减排序
func main() {
	p := make(FreqList, 10)
	for i := 0; i < 10; i++ {
		p[i].key = 10 - i
		p[i].value = (10 - i) * 2
	}
	fmt.Println(p)
	sort.Sort(p)
	fmt.Println(p)
}
