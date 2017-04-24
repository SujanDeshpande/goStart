package Subscriber

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
	"time"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func Subscribe(host string,username string, password string) {
	opts := MQTT.NewClientOptions().AddBroker(host)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetClientID("go-sub")
	opts.SetDefaultPublishHandler(f)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}


	go subscribe(c, "device.checkin")
	go subscribe(c, "device.enroll")
	go subscribe(c, "device.compliant")
	go subscribe(c, "device.noncompliant")
	go subscribe(c, "device.retired")

	fmt.Println("Waiting subscriber:")
	time.Sleep(30 * time.Second)
	fmt.Println("Waiting subscriber over:")
}

func subscribe(c MQTT.Client, topicName string) {

	if token := c.Subscribe(topicName, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	fmt.Println("Waiting now:" + topicName)
	time.Sleep(10 * time.Second)
	fmt.Println("Waiting over:")


	if token := c.Unsubscribe("device/checkin"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	c.Disconnect(20000)
}
