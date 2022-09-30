package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {

	var natsConn *nats.Conn
	if nc, err := nats.Connect(
		"ws://webstream:salasana@localhost:8080",
	); err != nil {
		log.Fatalln("nats.Connect", err)
	} else {
		natsConn = nc
	}

	var jetstream nats.JetStreamContext

	for {
		nats.JetStream

	}

}
