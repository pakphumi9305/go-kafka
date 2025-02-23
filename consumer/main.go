package main

import (
	"fmt"

	"gopkg.in/Shopify/sarama.v1"
)

func main() {
	servers := []string{"localhost:9092"}
	consumer, err := sarama.NewConsumer(servers, nil)

	if err != nil {
		panic(err)
	}
	//finishe consumer should be close
	defer consumer.Close()

	//topic = hello , consume partition 0 , offset is index that you can set to get data
	partitionConsumer, err := consumer.ConsumePartition("hello", 0, sarama.OffsetNewest)

	if err != nil {
		panic(err)
	}

	defer partitionConsumer.Close()

	fmt.Println("Consumer start.")

	//loop listening message
	for {
		select {
		case err := <-partitionConsumer.Errors():
			fmt.Println(err)
		case msg := <-partitionConsumer.Messages():
			fmt.Println(string(msg.Value))
		}
	}

}
