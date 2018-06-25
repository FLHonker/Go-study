package main

import (
	"net/rpc"
	"log"
	"os"
	"fmt"
	"errors"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith)Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t * Arith)Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

const serviceaddr = "127.0.0.1"

func main() {
	client, err := rpc.DialHTTP("tcp", serviceaddr + ":8811")
	if err != nil {
		log.Fatal("dialing:", err)
		os.Exit(1)
	}
	args := &Args{7, 8}
	var reply int
	//同步调用rpc
	err = client.Call("Arith.Multipl", args, &reply)
	if err != nil {
		log.Fatal("Arith erroe:", err)
	}
	fmt.Println("Arith:%d*%d=%d", args.A, args.B, reply)

	//异步调用rpc
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, &quotient, nil)
	replyCall := <-divCall.Done

	if replyCall != nil {
		fmt.Println("Arith:%d/%d=%d...%d", args.A, args.B, quotient.Quo, quotient.Rem)
	}

}