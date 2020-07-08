package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/dbr65/go-kafka-avro"
)

var kafkaServers = []string{"localhost:9092"}
var schemaRegistryServers = []string{"http://localhost:8081"}
var topic = "test"

func main() {
	var n int
	schema := `{
				"type": "record",
				"name": "Example",
				"fields": [
					{
						"name": "Number",
						"doc": "Phone number inside the national network. Length between 4-14",
						"type":  {
							  "type": "string",
							  "logicalType": "validated-string",
							  "pattern": "^[\\d]{4,14}$"
						}
					}
				]
			   }`
	producer, err := kafka.NewAvroProducer(kafkaServers, schemaRegistryServers)
	if err != nil {
		fmt.Printf("Could not create avro producer: %s", err)
	}
	flag.IntVar(&n, "n", 1, "number")
	flag.Parse()
	for i := 0; i < n; i++ {
		fmt.Println(i)
		addMsg(producer, schema)
	}
}

func addMsg(producer *kafka.AvroProducer, schema string) {
	value := `{
		"Number": "34675777777999999999999"
	}`
	key := time.Now().String()
	err := producer.Add(topic, schema, []byte(key), []byte(value))
	fmt.Println(key)
	if err != nil {
		fmt.Printf("Could not add a msg: %s", err)
	}
}
