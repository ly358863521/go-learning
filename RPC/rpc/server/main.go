package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

//计算乘积
func (t *Arith) Multiply(args *Args, reply *int) error {
	time.Sleep(time.Second * 3) //睡三秒，同步调用会等待，异步会先往下执行
	*reply = args.A * args.B
	return nil
}

//计算商和余数
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	time.Sleep(time.Second * 3)
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	//创建对象
	arith := new(Arith)
	//rpc服务注册了一个arith对象 公开方法供客户端调用
	rpc.Register(arith)
	//指定rpc的传输协议 这里采用http协议作为rpc调用的载体 也可以用rpc.ServeConn处理单个连接请求
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error", e)
	}
	go http.Serve(l, nil)
	os.Stdin.Read(make([]byte, 1))
}
