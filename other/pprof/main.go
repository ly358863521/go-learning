package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // 为什么用_ , 在讲解http包时有解释。
)

func myfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}

func main() {
	http.HandleFunc("/", myfunc)
	http.ListenAndServe(":8080", nil)
}

// 访问http://localhost:8080/debug/pprof/
