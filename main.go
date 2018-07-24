package main

import (
	"github.com/taybartski/log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	m := new(Mining)
	c := new(Client)
	server := rpc.NewServer()
	server.Register(m)
	server.Register(c)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Error("listen error: %v", e)
	}
	go test()
	for {
		if conn, err := listener.Accept(); err != nil {
			log.Error("accept error: %v", err.Error())
		} else {
			log.Info("new connection established")
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}

// type Args struct {
// X, Y int
// }

// type Calculator struct{}

// func (t *Calculator) Add(args *Args, reply *int) error {
// *reply = args.X + args.Y
// return nil
// }
