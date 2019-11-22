package main

import "fmt"

type Trie struct {
	isword   bool
	children map[int32]*Trie
}

func Constructor() *Trie {
	return &Trie{isword: false, children: make(map[int32]*Trie)}
}

func main() {
	a := Constructor()
	fmt.Println(a.children[1] == nil)
	a.children[1] = Constructor()
}
