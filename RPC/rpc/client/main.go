package main

import (
    "fmt"
    "log"
    "net/rpc"
    "time"
)

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

func main() {
    //调用rpc服务端提供的方法之前，先与rpc服务端建立连接
    client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
    if err != nil {
        log.Fatal("dialHttp error", err)
        return
    }
    //同步调用服务端提供的方法

    args := &Args{7, 8}
    var reply int
    //可以查看源码 其实Call同步调用是用异步调用实现的。后续再详细学习
    err = client.Call("Arith.Multiply", args, &reply) //这里会阻塞三秒

    if err != nil {
        log.Fatal("call arith.Multiply error", err)
    }
    fmt.Printf("Arith:%d*%d=%d\n", args.A, args.B, reply)

    //异步调用

    quo := Quotient{}

    divCall := client.Go("Arith.Divide", args, &quo, nil)

    //使用select模型监听通道有数据时执行，否则执行后续程序
    for {
        select {
        case <-divCall.Done:
            fmt.Printf("商是%d,余数是%d\n", quo.Quo, quo.Rem)
        default:
            fmt.Println("继续向下执行....")
            time.Sleep(time.Second * 1)
        }
    }

}