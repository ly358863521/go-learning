## 发送方

### 主进程

发送连接请求

### 发送进程

1. 默认状态下，设置超时计时后，发送滑动窗口内序号的数据包，发送完毕堵塞等待。

2. 等待超时信号，超时情况下，开启新的发送进程，滑动窗口不变，结束当前发送进程。

3. 等待ack信号，收到确认时，传出当前待发送的数据包序号，结束当前进程。

4. ```go
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
   ```

### 接收进程

1. 接收返回的数据包，解析确认ack，判断ack是否在当前滑动窗口中。
2. 发送ack信号给发送进程，获取当前待发送数据包序号同时结束掉当前发送进程。
3. 移动滑动窗口到当前ack序号后，并开启新的发送进程，传入当前待发送序号。

## 接收方

### 主进程

建立监听

### 接收进程

1. 接收当前连接发来的数据包，解析序号，序号与待接收序号相等就接收，否则丢弃。
2. 接收数据包后，发送确认ack给发送方。