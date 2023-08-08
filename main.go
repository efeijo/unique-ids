package main

import (
	"encoding/json"
	"log"
	"net/rpc"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	n.Handle("generate", func(msg maelstrom.Message) error {
		var body map[string]any
		err := json.Unmarshal(msg.Body, &body)
		if err != nil {
			log.Fatal("error unmarshilling", err)
			return err
		}

		var id int64
		err = client.Call("IdGen.GenerateIds", struct{}{}, &id)
		if err != nil {
			time.Sleep(2 * time.Second)
			client.Call("IdGen.GenerateIds", struct{}{}, &id)
		}
		body["type"] = "generate_ok"
		body["id"] = id

		err = n.Reply(msg, body)
		if err != nil {
			client.Call("IdGen.ErrorDecrement", struct{}{}, struct{}{})
		}
		return nil
	})

	if err := n.Run(); err != nil {
		log.Fatal("erro running node", err)
	}

}

type MyUniqueIdService struct {
	id int64
	Ch chan int64
}

func NewMyUniqueService() {

}
