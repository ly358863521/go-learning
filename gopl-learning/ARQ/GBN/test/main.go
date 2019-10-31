package main

import (
	"fmt"
	win "foo/windows"
)

func main() {
	var w win.Window
	w.Init(10)
	fmt.Println(w.WinSize)
	for i := 0; i < w.WinSize; i++ {
		w.Timer[i] = win.Timer(i)
		w.Status[i] = true
	}
	fmt.Println(w.Timer, w.Status)
	w.Move(4)
	fmt.Println(w.Timer, w.Status)
	fmt.Println(w.Find(5))
}
