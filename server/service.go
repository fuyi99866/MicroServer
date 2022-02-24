package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/server"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int

var (
	addr = flag.String("addr","localhost:8972","server address")
)

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()
	s := server.NewServer()
	s.RegisterName("Arith", new(Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}

}
