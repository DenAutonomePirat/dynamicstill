package main

import (
	"fmt"
	"gopkg.in/redis.v5"
)

func main() {
	client := redis.NewClient(&redis.Options{
		//Network:  "unix",
		//Addr:     "/tmp/redis.sock",
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pubsub, err := client.Subscribe("test")
	if err != nil {
		panic(err)
	}
	defer pubsub.Close()

	for {

		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			panic(err)
		}
		fmt.Println(msg.Channel, msg.Payload)
	}
}
