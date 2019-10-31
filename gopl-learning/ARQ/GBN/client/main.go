package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

type win struct {
	start uint64
	end   uint64
}

func (p *win) init(wind_size int) {
	p.start = 0
	p.end = uint64(wind_size) - 1
}

func (p *win) move(n int) {
	p.start += uint64(n)
	p.end += uint64(n)
}
func send(w win, curseq uint64, conn *net.UDPConn) {
	if w.end > totalPacket {
		w.end = totalPacket
	}
	timeout := time.After(5 * time.Second)
loop:
	for {
		select {
		case <-timeout:
			fmt.Println("超时重传")
			go send(w, w.start, conn)
			break loop
		case <-ack:
			fmt.Println("当前待发送序列号为:", curseq)
			seq <- curseq
			break loop
		default:
			if curseq < uint64(w.end) {
				bs := make([]byte, 1032)
				binary.BigEndian.PutUint64(bs, curseq+1)
				conn.Write(bs)
				curseq++
			}
		}
	}
}
func recv(w win, conn *net.UDPConn) {
	for {
		data := make([]byte, 8)
		_, err := conn.Read(data)
		if err != nil {
			fmt.Println("failed to read UDP msg because of ", err)
			os.Exit(1)
		}
		curack := binary.BigEndian.Uint64(data)
		fmt.Println("收到确认ack：", curack)
		if curack+1 == totalPacket {
			bs := make([]byte, 8)
			binary.BigEndian.PutUint64(bs, 0)
			conn.Write(bs)
			ack <- 0
			done <- struct{}{}
			return
		}
		ack <- curack
		w.move(int(curack-w.start) + 1)
		go send(w, <-seq, conn)
	}
}

var (
	host                      = flag.String("host", "localhost", "host")
	port                      = flag.String("port", "8080", "port")
	wind_size                 = 10
	seq_size                  = 10
	ack         chan uint64   = make(chan uint64, 1)
	seq         chan uint64   = make(chan uint64, 1)
	curAck      int           //当前等待确认的ack
	totalSeq    int           //收到的包的总数
	totalPacket uint64        //需要发送的包总数
	done        chan struct{} = make(chan struct{})
)

//go run timeclient.go -host time.nist.gov
func main() {
	totalPacket = 20
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
	var w win
	w.init(10)
	go send(w, 0, conn)
	go recv(w, conn)
	// bs := make([]byte, 8)
	// binary.BigEndian.PutUint64(bs, 123)
	// fmt.Println(bs)
	// // bs := []byte(strconv.Itoa(123))
	// // fmt.Println(bs)
	// _, err = conn.Write(bs)
	// if err != nil {
	// 	fmt.Println("failed:", err)
	// 	os.Exit(1)
	// }
	// data := make([]byte, 4)
	// _, err = conn.Read(data)
	// if err != nil {
	// 	fmt.Println("failed to read UDP msg because of ", err)
	// 	os.Exit(1)
	// }
	// t := binary.BigEndian.Uint32(data)
	// fmt.Println(time.Unix(int64(t), 0).String())
	<-done
}
