package main

import (
	"fmt"
)

func main() {
	abort := make(chan struct{})
	done := make(chan struct{})
	go func() {
	loop:
		for i := 0; i < 10; i++ {
			select {
			case <-abort:
				fmt.Println("break loop")
				break loop
			default:
				fmt.Println(i)
				if i == 7 {
					go func() {
						abort <- struct{}{}
					}()
				}
			}
		}
		fmt.Println("done")
		done <- struct{}{}
	}()
	<-done
}
