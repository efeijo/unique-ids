package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type IdGen struct {
	id        int64
	idChannel chan int64
}

func NewIdGen() *IdGen {
	ch := make(chan int64)
	ig := &IdGen{
		idChannel: ch,
		id:        0,
	}
	go func() {
		for {
			ch <- ig.id
			ig.id++
		}
	}()

	return ig
}

func (ig *IdGen) GenerateIds(args struct{}, reply *int64) error {
	r := <-ig.idChannel
	// simple bug man
	// 	reply = &r :/
	*reply = r
	return nil
}

func main() {
	idGen := NewIdGen()
	rpc.Register(idGen)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil)
}
