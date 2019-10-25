package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

var clients  = make(map[string]bool)

type Pxy struct {}

func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)

	transport :=  http.DefaultTransport

	// step 1
	outReq := new(http.Request)
	*outReq = *req // this only does shallow copies of maps

	clientIP, _, err := net.SplitHostPort(req.RemoteAddr)
	if err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
		//X-Forwarded-For: client, proxy1, proxy2
	}

	
	if _,ok := clients[clientIP];!ok{
		outReq.Host = "127.0.0.1:12345"
		fmt.Println(outReq.Host)
		res, err := transport.RoundTrip(outReq)
		if err != nil {
			rw.WriteHeader(http.StatusBadGateway)
			return
		}
		for key, value := range res.Header {
			for _, v := range value {
				rw.Header().Add(key, v)
			}
		}
		rw.WriteHeader(res.StatusCode)
		io.Copy(rw, res.Body)
		res.Body.Close()
	}else{

		fmt.Printf("send request %s %s %s\n", outReq.Method, outReq.Host, outReq.RemoteAddr)
		res, err := transport.RoundTrip(outReq)
		if err != nil {
			rw.WriteHeader(http.StatusBadGateway)
			return
		}

		// step 3
		for key, value := range res.Header {
			for _, v := range value {
				rw.Header().Add(key, v)
			}
		}
		rw.WriteHeader(res.StatusCode)
		io.Copy(rw, res.Body)
		res.Body.Close()
	}
	// step 2
	clients[clientIP] = true
}

func main() {
	fmt.Println("Serve on :8080")
	http.Handle("/", &Pxy{})
	http.ListenAndServe("0.0.0.0:8080", nil)
}