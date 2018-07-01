package rpcClient

/*
Usage: httpRPC_Client localhost
 */

import (
	"os"
	"fmt"
	"net/rpc"
	"log"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func httpRPC_Client() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], " server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]
	client, err := rpc.DialHTTP("tcp", serverAddress + ":8811")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//Synchronous
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatalf("arith error:", err)
	}
	fmt.Printf("Arith %d*%d=%d\n", args.A, args.B, reply)

	var quo Quotient
	err = client.Call("Arith.Divide", args, &quo)
	if err != nil {
		log.Fatalf("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d ... %d\n", args.A, args.B, quo.Quo, quo.Rem)

}