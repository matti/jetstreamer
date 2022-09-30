package main

import (
	"log"
	"time"

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
	if js, err := natsConn.JetStream(nats.PublishAsyncMaxPending(1)); err != nil {
		log.Fatalln("natsConn.JetStream", err)
	} else {
		jetstream = js
	}

	// Create a Stream
	js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.*"},
	})

	for {
		log.Println("publishing")

		if pa, err := jetstream.Publish("hello", []byte("world")); err != nil {
			log.Println("jetstream.Publish", err)
		} else {
			log.Println("published", pa.Sequence, pa.Domain, pa.Stream, pa.Duplicate)
		}

		time.Sleep(time.Second * 1)
	}

}
