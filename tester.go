package main

import (
	"github.com/taybartski/log"
	"net"
	"net/rpc/jsonrpc"
	"time"
)

type Args struct {
	X, Y int
}

func test() {
	time.Sleep(1)
	client, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Error("dialing: %v", err)
	}
	// Synchronous call
	args := &Args{7, 8}
	var reply int
	c := jsonrpc.NewClient(client)
	err = c.Call("Calculator.Add", args, &reply)
	if err != nil {
		log.Error("arith error: %v", err)
	}
	log.Info("Result: %v", reply)
}
