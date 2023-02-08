package kafka

import (
	"log"
	"encoding/json"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	route2 "eulevo.com/delivey/simulator/application/route"
	"eulevo.com/delivey/simulator/infra/kafka"

)

//{"clientId":"1", "routeId":"1"}
//{"clientId":"2", "routeId":"2"}
func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publisher(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}