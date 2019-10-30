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
	host               = flag.String("host", "", "host")
	port               = flag.String("port", "8080", "port")
	wind_size          = 10
	seq_size           = 10
	endrecv            = -1
	ack         []bool = make([]bool, seq_size) //收到ack情况
	curAck      int                             //当前等待确认的ack
	totalSeq    int                             //收到的包的总数
	totalPacket int                             //需要发送的包总数
)

func main() {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	for {
		data := make([]byte, 8)
		n, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println("failed to read UDP msg because of ", err.Error())
			return
		}
		// fmt.Println(data)
		seq := binary.BigEndian.Uint64(data)
		// seq, err := strconv.Atoi(string(data[:n]))
		fmt.Println(seq)
		daytime := time.Now().Unix()
		fmt.Println(n, remoteAddr)
		b := make([]byte, 4)
		binary.BigEndian.PutUint32(b, uint32(daytime))
		conn.WriteToUDP(b, remoteAddr)

	}
	// daytime := time.Now().Unix()
	// fmt.Println(n, remoteAddr)
	// b := make([]byte, 4)
	// binary.BigEndian.PutUint32(b, uint32(daytime))
	// conn.WriteToUDP(b, remoteAddr)
}
