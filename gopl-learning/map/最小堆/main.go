package main

type Freq struct {
	key   int
	value int
}
type FreqList []Freq
func topKFrequent(nums []int, k int) []int {
	dict := make(map[int]int)
	for _, i := range nums {
		dict[i]++
	}
	Frequency := []*Freq{}
	i:=0
	for k,v:=range dict{
		Frequency[i].key = k
		Frequency[i++].value = v
	}
	
}
func main() {

}
