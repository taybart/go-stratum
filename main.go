package main

import (
	"github.com/taybartski/log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var l log.Log

/*
SPEC: https://en.bitcoin.it/wiki/Stratum_mining_protocol
Methods (client to server)
mining.authorize("username", "password")
mining.capabilities({"notify":[], "set_difficulty":{}, "set_goal":{}, "suggested_target": "hex target"})
mining.extranonce.subscribe()
mining.get_transactions("job id")
mining.submit("username", "job id", "ExtraNonce2", "nTime", "nOnce")
mining.subscribe("user agent/version", "extranonce1")
mining.suggest_difficulty(preferred share difficulty Number)
mining.suggest_target("full hex share target")
...
Methods (server to client)
client.get_version()
client.reconnect("hostname", port, waittime)
client.show_message("human-readable message")
mining.notify(...)
mining.set_difficulty(difficulty)
mining.set_extranonce("extranonce1", extranonce2_size)
mining.set_goal("goal name", {"malgo": "SHA256d", ...})
*/

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
