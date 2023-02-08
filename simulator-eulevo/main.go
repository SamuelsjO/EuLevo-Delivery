package main

import (
	"log"
	"fmt"

    "github.com/joho/godotenv"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	kafka2 "eulevo.com/delivey/simulator/application/kafka"
	"eulevo.com/delivey/simulator/infra/kafka"

)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main(){
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}

}