package main

import (
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/Shopify/sarama.v1"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	//get config in envi first
	viper.AutomaticEnv()
	//ใน env sh ใช้ . ไม่ได้ต้อง replace
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka"))

	if err != nil {
		panic(err)
	}
	// consumer.Consume(”)
}

// func main() {
// 	servers := []string{"localhost:9092"}
// 	consumer, err := sarama.NewConsumer(servers, nil)

// 	if err != nil {
// 		panic(err)
// 	}
// 	//finishe consumer should be close
// 	defer consumer.Close()

// 	//topic = hello , consume partition 0 , offset is index that you can set to get data
// 	partitionConsumer, err := consumer.ConsumePartition("hello", 0, sarama.OffsetNewest)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer partitionConsumer.Close()

// 	fmt.Println("Consumer start.")

// 	//loop listening message
// 	for {
// 		select {
// 		case err := <-partitionConsumer.Errors():
// 			fmt.Println(err)
// 		case msg := <-partitionConsumer.Messages():
// 			fmt.Println(string(msg.Value))
// 		}
// 	}

// }
