package main

import (
	"ActiveMQ/Producer"
	"fmt"
	"ActiveMQ/Subscriber"
	"time"
)

func main() {
	host := "tcp://52.3.243.201:1883"
	username := "admin@foo.com"
	password := "Mi4man11"

	go Subscriber.Subscribe(host,username,password)
	fmt.Println("Waiting between:")
	time.Sleep(5 * time.Second)
	fmt.Println("Waiting between over:")
	go Producer.Produce(host,username,password)
	fmt.Println("Waiting main:")
	time.Sleep(50 * time.Second)
	fmt.Println("Waiting main over:")

}
