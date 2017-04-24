package Producer

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	"os"
	"time"
)

func Produce(host string,username string, password string) {
	opts := MQTT.NewClientOptions().AddBroker(host)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetClientID("go-pub")

	text := fmt.Sprintf("this is msg")

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	go publish(c, "42000/device.checkin", text)
	go publish(c, "42000/device.enroll", text)
	go publish(c, "42000/device.compliant", text)
	go publish(c, "42000/device.noncompliant", text)
	go publish(c, "42000/device.retired", text)

	fmt.Println("Waiting producer now:")
	time.Sleep(10 * time.Second)
	fmt.Println("Waiting producer over:")
	c.Disconnect(20000)
}

func readDevice() {
	file, e := ioutil.ReadFile("./device.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	jsonString := string(file)
	fmt.Printf("%s\n", jsonString)
	jsonBytes := []byte(jsonString)
	device := new(Device)
	err := json.Unmarshal(jsonBytes, &device)
	if err != nil {
		fmt.Println("Error")
	}

	fmt.Println(device)
}

func publish(c MQTT.Client, topicName string, text string) {

	token := c.Publish(topicName, 1, false, text)
	token.Wait()
	fmt.Println(topicName)

}

type Device struct {
	Timestamp int64  `json:"timestamp"`
	EventType string `json:"eventType"`
	Devices   []struct {
		Compliant        bool   `json:"compliant"`
		Status           string `json:"status"`
		LastCheckInTime  int64  `json:"lastCheckInTime"`
		RegistrationTime int64  `json:"registrationTime"`
		Identifier       string `json:"identifier"`
		MacAddress       string `json:"macAddress"`
		Manufacturer     string `json:"manufacturer"`
		Model            string `json:"model"`
		Os               string `json:"os"`
		OsVersion        string `json:"osVersion"`
		SerialNumber     string `json:"serialNumber"`
		UserID           string `json:"userId"`
		UserUUID         string `json:"userUuid"`
	} `json:"devices"`
}
