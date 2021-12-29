package main

import (
	"encoding/json"
	"fmt"
	"github.com/Saimunyz/L0/internal/model"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"os"
)

const (
	clientID  = "publisher"
	clusterID = "test-cluster"
	channel   = "session"
)

func main() {

	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		fmt.Println(err)
	}
	defer sc.Close()

	jsonFIle, err := os.Open("data.json")
	if err != nil {
		log.Printf("error while openning file %v", err)
	}
	defer jsonFIle.Close()

	byteValue, _ := ioutil.ReadAll(jsonFIle)

	var data []model.Data

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Printf("error while parsing file %v", err)
	}

	data = append(data, model.Data{})

	for _, value := range data {
		bytesValue, _ := json.Marshal(value)

		err = sc.Publish(channel, bytesValue)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = sc.Publish(channel, []byte("Some wrong data"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("6 messages have been sent: 4 messages are normal and 2 bad")
}
