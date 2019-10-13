// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages

	//8.12:使broadcaster能够将arrival事件通知当前所有的客户端。为了达成这个目的，你需要有一个客户端的集合，并且在entering和leaving的channel中记录客户端的名字。
	clientMap = make(map[string]string)
	oldkey    = int64(1001)
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {

				//8.15:如果一个客户端没有及时地读取数据可能会导致所有的客户端被阻塞。修改broadcaster来跳过一条消息，而不是等待这个客户端一直到其准备好写
				if !clients[cli] {
					continue
				}
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()

	//8.14:修改聊天服务器的网络协议这样每一个客户端就可以在entering时可以提供它们的名字。将消息前缀由之前的网络地址改为这个名字。
	if len(clientMap[who]) <= 0 {
		str := "num" + fmt.Sprint(oldkey)
		clientMap[who] = str
		oldkey++
	}
	ch <- "You are " + clientMap[who]
	messages <- clientMap[who] + " has arrived"
	entering <- ch //ch->entering->client (clent = ch )-> msg->client <->msg->ch

	input := bufio.NewScanner(conn)

	// 8.13:使聊天服务器能够断开空闲的客户端连接，比如最近五分钟之后没有发送任何消息的那些客户端

	abort := make(chan string)
	go func() {
		for {
			select {
			case <-time.After(5 * time.Minute):
				conn.Close()
			case str := <-abort:
				fmt.Println("str:", str)
				messages <- str
			}
		}
	}()
	for input.Scan() {
		str := input.Text()
		if str == "exit" {
			break
		}
		// if len(str) > 0 {
		// 	abort <- clientMap[who] + ":" + str
		// 	fmt.Println("abort!")
		// }
		abort <- clientMap[who] + ":" + str
		// messages <- clientMap[who] + ": " + str
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- clientMap[who] + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		//只循环但没有取
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
