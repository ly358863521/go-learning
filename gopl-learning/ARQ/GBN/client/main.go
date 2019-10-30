package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var (
	host               = flag.String("host", "localhost", "host")
	port               = flag.String("port", "8080", "port")
	wind_size          = 10
	seq_size           = 10
	ack         []bool = make([]bool, seq_size) //收到ack情况
	curSeq      int                             //当前数据包序号
	curAck      int                             //当前等待确认的ack
	totalSeq    int                             //收到的包的总数
	totalPacket int                             //需要发送的包总数
)

//go run timeclient.go -host time.nist.gov
func main() {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Can't dial: ", err)
		os.Exit(1)
	}
	fmt.Println("connected!")
	defer conn.Close()
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, 123)
	fmt.Println(bs)
	// bs := []byte(strconv.Itoa(123))
	// fmt.Println(bs)
	_, err = conn.Write(bs)
	if err != nil {
		fmt.Println("failed:", err)
		os.Exit(1)
	}
	data := make([]byte, 4)
	_, err = conn.Read(data)
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err)
		os.Exit(1)
	}
	t := binary.BigEndian.Uint32(data)
	fmt.Println(time.Unix(int64(t), 0).String())
	os.Exit(0)
}
