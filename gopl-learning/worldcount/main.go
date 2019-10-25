package main
 
import (
        "bufio"
        "fmt"
        "strings"
)
 
func main() {
 
        var c WordsCounter
        fmt.Fprintf(&c, "hello world 123")
		fmt.Println(c) //输出 3
		// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
		// 	p := newPrinter()
		// 	p.doPrintf(format, a)
		// 	n, err = w.Write(p.buf)
		// 	p.free()
		// 	return
		// }
}

type ByteCounter int
 
func (c *ByteCounter) Write(p []byte) (int, error) {
        *c += ByteCounter(len(p)) // convert int to ByteCounter
        return len(p), nil
}
 
//定义类型
type WordsCounter int
 
//满足相同接口的类型
func (w *WordsCounter) Write(p []byte) (int, error) {
        //分隔字符串
        s := strings.NewReader(string(p))
        bs := bufio.NewScanner(s)
        bs.Split(bufio.ScanWords)
        sum := 0
        for bs.Scan() {
                sum++
        }  
        *w = WordsCounter(sum)
        return sum, nil
}