package main

import (
	"crypto/tls"
	"log"
	"time"
	"crypto/rand"
	"net"
	"io"
)

func main() {
	cert, err := tls.LoadX509KeyPair("rui.crt", "rui.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates:[]tls.Certificate{cert}}
	config.Time = time.Now
	config.Rand = rand.Reader

	service := "127.0.0.1:8811"
	listener, err := tls.Listen("tcp", service, &config)
	if err != nil {
		log.Fatalf("server: listen: %s", err)
	}
	log.Println("server: listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept: %s", err)
			break
		}
		log.Printf("server: accepted from %s", conn.RemoteAddr())
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 512)
	for {
		log.Println("server: conn: waiting")
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("server: conn: read: %s", err)
			}
			break
		}
		log.Printf("server: conn: echo %q\n", string(buf[:n]))
		if err != nil {
			log.Printf("server: write: %s", err)
			break
		}
		log.Println("erver: conn: closed")
	}
}