package main

import (
	"net"
	"fmt"
	"time"
	"os"
)

/*
func ListenTCP(net string, laddr *TCPAddr) (l *TCPListener, err os.Error)
func (l *TCPListener) Accept() (c Conn, err os.Error)
*/

// 服务端server

func main() {
	service := ":8811"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)
	fmt.Println("Servering ...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)	//多线程处理
	}
}

// 处理client请求
func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Println("A new client conn...")
	daytime := time.Now().String()
	conn.Write([]byte(daytime))		//don't care about return value
	// We're finished with this client

}

// 错误处理
func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}