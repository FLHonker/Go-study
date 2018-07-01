package main

import (
	"os"
	"fmt"
	"net"
	"io/ioutil"
)

/*
func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)
	net参数是	"tcp4"、"tcp6"、"tcp"中的任意一个,分别表示TCP(IPv4-only)、TCP(IPv6-only)或者
	TCP(IPv4,IPv6的任意一个)
	laddr表示本机地址,一般设置为nil
	raddr表示远程的服务地址
 */

// 客户端client
// test web: baidu.com -- 14.215.177.39:80
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkErr(err)
	_, err = conn.Write([]byte("Head / HTTP/1.0\r\n\r\n"))
	checkErr(err)
	result, err := ioutil.ReadAll(conn)
	checkErr(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}