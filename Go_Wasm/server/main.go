package main

import (
    "fmt"
	"net/http"
	"io"
	"os"
)

func FileHandle(w http.ResponseWriter, r *http.Request) {
    // 这里一定要记得 r.ParseMultipartForm(), 否则 r.MultipartForm 是空的
    // 调用 r.FormFile() 的时候会自动执行 r.ParseMultipartForm()
    r.ParseMultipartForm(32 << 20) 
    // 写明缓冲的大小。如果超过缓冲，文件内容会被放在临时目录中，而不是内存。过大可能较多占用内存，过小可能增加硬盘 I/O
    // FormFile() 时调用 ParseMultipartForm() 使用的大小是 32 << 20，32MB
    file, fileHeader, err := r.FormFile("uploadfile") // file 是上传表单域的名字
    if err != nil {
        fmt.Println("get upload file fail:", err)
        w.WriteHeader(500)
        return
    }
    defer file.Close() // 此时上传内容的 IO 已经打开，需要手动关闭！！

    // fileHeader 有一些文件的基本信息
    fmt.Println(fileHeader.Header.Get("Content-Type"))
    fmt.Println(fileHeader.Filename)
    fmt.Println(fileHeader.Size)
    // 打开目标地址，把上传的内容存进去
    f, err := os.OpenFile(fileHeader.Filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        fmt.Println("save upload file fail:", err)
        w.WriteHeader(500)
        return
    }

    defer f.Close()
    if _, err = io.Copy(f, file); err != nil {
        fmt.Println("save upload file fail:", err)
        w.WriteHeader(500)
        return
    }
    w.Write([]byte("upload file:" + fileHeader.Filename ))
}

func main(){
    fs := http.FileServer(http.Dir("."))
    http.Handle("/", fs)
	http.HandleFunc("/upload", FileHandle)
    if err := http.ListenAndServe(":12345",nil); err != nil{
        fmt.Println("start http server fail:",err)
    }
}