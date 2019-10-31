package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	host    = flag.String("host", "", "host")
	port    = flag.String("port", "8080", "port")
	endrecv = 1
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
		data := make([]byte, 1032)
		_, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println("failed to read UDP msg because of ", err.Error())
			return
		}
		// fmt.Println(data)
		seq := binary.BigEndian.Uint64(data[:8])
		// seq, err := strconv.Atoi(string(data[:n]))
		if int(seq) == endrecv {
			fmt.Println("收到数据包序列号：", seq-1)
			ack := make([]byte, 8)
			binary.BigEndian.PutUint64(ack, seq-1)
			conn.WriteToUDP(ack, remoteAddr)
			endrecv++
		}
		if int(seq) == 0 {
			fmt.Println("收到序列号0，接收完毕！")
			endrecv = 1
		}
		// daytime := time.Now().Unix()
		// fmt.Println(n, remoteAddr)
		// b := make([]byte, 4)
		// binary.BigEndian.PutUint32(b, uint32(daytime))
		// conn.WriteToUDP(b, remoteAddr)

	}
	// daytime := time.Now().Unix()
	// fmt.Println(n, remoteAddr)
	// b := make([]byte, 4)
	// binary.BigEndian.PutUint32(b, uint32(daytime))
	// conn.WriteToUDP(b, remoteAddr)
}
