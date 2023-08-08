package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct{}

type IdGen struct {
	id        int64
	idChannel chan int64
	opChannel chan bool
}

func NewIdGen() *IdGen {
	ch := make(chan int64)
	opCh := make(chan bool)
	ig := &IdGen{
		idChannel: ch,
		id:        0,
		opChannel: opCh,
	}
	go func() {
		for {
			op := <-ig.opChannel
			if op {
				ch <- ig.id
				ig.id++
			} else {
				ig.id--
			}
		}
	}()

	return ig
}

func (ig *IdGen) GenerateIds(args struct{}, reply *int64) error {
	ig.opChannel <- true
	r := <-ig.idChannel
	reply = &r
	fmt.Println("Generated id:", *reply)
	return nil
}

func (ig *IdGen) ErrorDecrement(args struct{}, reply struct{}) error {
	ig.opChannel <- false
	fmt.Println("decremented")
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
