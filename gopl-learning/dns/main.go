package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

var data = make(map[string]string)

func Init(){
	f, err := ioutil.ReadFile("ch1/server1/data.json")
	if err!=nil{
		fmt.Println(err)
	}
	json.Unmarshal(f,&data)
}
func writeback(){
	str,err:= json.Marshal(data)
	if err != nil{
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	if ioutil.WriteFile("ch1/server1/data.json",str,0777) == nil {
        fmt.Println("suceess!")
	}
}

func main(){

	Init()
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	var res string
	switch r.URL.Path {
	case "/add":
		name := r.URL.Query().Get("name")
		address :=r.URL.Query().Get("address")
		data[name] = address
		writeback()
		res = "Success!"

	case "/get":
		name := r.URL.Query().Get("name")
		res = data[name]
	}
	fmt.Fprintf(w,res)
}

//!-
