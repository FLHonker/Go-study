package main

import (
	"errors"
	"net/rpc"
	"net"
	"fmt"
	"os"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divided by zero.")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8811")
	checkErr(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)		//important!
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}