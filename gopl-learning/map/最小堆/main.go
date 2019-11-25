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
	Frequency := make(FreqList, len(dict))
	i := 0
	for k, v := range dict {
		Frequency[i].key = k
		Frequency[i].value = v
		i++
	}
	heap := Frequency[:k]
	buildheap(heap, k)
	for i := k; i < len(Frequency); i++ {
		if Frequency[i].value > heap[0].value {
			heap = heap[1:]
			heap = append(heap, Frequency[i])
			buildheap(heap, k)
		}
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap[i].key
	}
	return res

}

func heapify(freq FreqList, n int, i int) {
	left, right := i*2+1, i*2+2 //左右节点
	min := i
	if left < n && freq[min].value > freq[left].value {
		min = left
	}
	if right < n && freq[min].value > freq[right].value {
		min = right
	}
	if min != i {
		freq[min], freq[i] = freq[i], freq[min]
		heapify(freq, n, min)
	}
}
func buildheap(freq FreqList, n int) {
	for i := n / 2; i >= 0; i-- {
		heapify(freq, n, i)
	}

}
