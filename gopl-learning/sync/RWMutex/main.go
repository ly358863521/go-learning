package main

import(
	"sync"
	"fmt"
	"time"
)
var mu sync.RWMutex
func read(i int){
	mu.Lock()
	//defer mu.Unlock()
	fmt.Println("readed",i)
}
func main(){
	//mu.RLock()
	for i :=0;i<10;i++{
		fmt.Println("start",i)
		go read(i)
		time.Sleep(1*time.Second)
	}
}