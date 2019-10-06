package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(url, filename string, filePath string) error {

	//打开文件句柄操作
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	r,w :=io.Pipe()
    //创建一个模拟的form中的一个选项,这个form项现在是空的
	// bodyBuf := &bytes.Buffer{}
	// bodyWriter := multipart.NewWriter(bodyBuf)
	bodyWriter := multipart.NewWriter(w)
	go func(){
		defer w.Close()
		defer bodyWriter.Close()
		fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
		if err != nil {
			fmt.Println("error writing to buffer")
			return
		}
		defer file.Close()
    	if _, err = io.Copy(fileWriter, file); err != nil {
			return
		}

		//fileWriter is io.Writer 
		// type Writer interface {
		// 	Write(p []byte) (n int, err error)
		// }
		bodyWriter.Close()
		params := map[string]string{
			"filename" : filename,
		}
		for key, val := range params {
			_ = bodyWriter.WriteField(key, val)
	}
	}()
	// bodyWriter := multipart.NewWriter(w)
	// //关键的一步操作, 设置文件的上传参数叫uploadfile, 文件名是filename,
	// //相当于现在还没选择文件, form项里选择文件的选项
	// fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	// if err != nil {
	// 	fmt.Println("error writing to buffer")
	// 	return err
	// }

	// //iocopy 这里相当于选择了文件,将文件放到form中
	// _, err = io.Copy(fileWriter, file)
	// if err != nil {
	// 	return err
	// }

    // //获取上传文件的类型,multipart/form-data; boundary=...
	// contentType := bodyWriter.FormDataContentType()
	// bodyWriter.Close()
	// params := map[string]string{
    //     "filename" : filename,
    // }
	// for key, val := range params {
	// 	_ = bodyWriter.WriteField(key, val)
	// }

    //发送post请求到服务端
	resp, err := http.Post(url,bodyWriter.FormDataContentType(), r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

// sample usage
func main() {
	url := "http://localhost:12345"
    filename := "test.txt"
	file := "client/test.txt" //上传的文件


	postFile(url, filename, file)
}
