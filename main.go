package main

import (
	"github.com/taybartski/log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var l log.Log

type Args struct {
	X, Y int
}

type Calculator struct{}

func (t *Calculator) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func main() {
	cal := new(Calculator)
	server := rpc.NewServer()
	server.Register(cal)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		l.Error("listen error: %v\n", e)
	}
	for {
		if conn, err := listener.Accept(); err != nil {
			l.Error("accept error: %v\n", err.Error())
		} else {
			l.Info("new connection established\n")
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}
