package main

import (
	"fmt"
	"os"
	"net"
	"time"
)

// 服务端server

func main() {
	service := ":8811"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime:= time.Now().String()
		conn.Write([]byte(daytime))		//don't care about return value
		conn.Close()		//We're finished with this client
	}
}

// 错误处理
func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}