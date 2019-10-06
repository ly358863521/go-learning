package main

import (
	"time"
	"fmt"
	"syscall/js"
	"io"
	"mime/multipart"
	"net/http"
	"io/ioutil"

)

var (
	r *io.PipeReader
	w *io.PipeWriter
	bodyWriter *multipart.Writer
	fileWriter io.Writer
)

var filename string
func t(this js.Value, params []js.Value) interface{} {
	value := time.Now().Format("2006-01-02 15:04:05")
	js.Global().Set("output",js.ValueOf(value))
	return js.ValueOf(value)
}

func receive(this js.Value, params []js.Value) interface{}{
	data := make([]byte,params[0].Length())
	js.CopyBytesToGo(data,params[0])

	var err error
	if fileWriter == nil{
		fileWriter, err = bodyWriter.CreateFormFile("uploadfile", filename)
	}
	if err != nil {
		fmt.Println("error writing to buffer")
		fmt.Println(err)
		return js.ValueOf(false)
	}
	if _, err = fileWriter.Write(data); err != nil {
		fmt.Println("error writing data")
		fmt.Println(err)
		return js.ValueOf(false)
	}

	return js.ValueOf(true)
}

func finishPost(this js.Value,params []js.Value)interface{}{
	bodyWriter.Close()
	w.Close()
	return nil
}


func postfile(this js.Value,params []js.Value)interface{}{
	url := params[0].String()
	filename = params[1].String()
	r,w =io.Pipe()
	bodyWriter = multipart.NewWriter(w)
	// bodyParams := map[string]string{
	// 	"filename" : filename,
	// }
	// for key, val := range bodyParams {
	// 	_ = bodyWriter.WriteField(key, val)
	// }
	defer w.Close()
	defer bodyWriter.Close()
	fmt.Println("to get data")
	resp, err := http.Post(url,bodyWriter.FormDataContentType(), r)
	fmt.Println("send post done")
	if err != nil {
		fmt.Println(err)
		return js.ValueOf(false)
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return js.ValueOf(false)
	}
	fmt.Println(resp.Status)
	//fmt.Println(string(resp_body))
	return js.ValueOf(string(resp_body))
}


func setcb(){
	js.Global().Set("time",js.FuncOf(t))
	js.Global().Set("receive",js.FuncOf(receive))
	js.Global().Set("finishPost",js.FuncOf(finishPost))
	js.Global().Set("postfile",js.FuncOf(func(this js.Value,params []js.Value)interface{}{
		go postfile(this,params);
		return nil;
	}))
}



func main() {
	
	fmt.Println("Hello, WebAssembly!!!!")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	setcb()
	select {}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}