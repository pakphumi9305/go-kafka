package main

import (
	"producer/controllers"
	"producer/services"
	"strings"

	"github.com/gofiber/fiber/v2"
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
	producer, err := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)

	if err != nil {
		panic(err)
	}

	defer producer.Close()

	eventProducer := services.NewEventProducer(producer)
	accountService := services.NewAccountService(eventProducer)
	accountController := controllers.NewAccountController(accountService)

	app := fiber.New()
	app.Post("/openAccount", accountController.OpenAccount)
	app.Post("/depositFund", accountController.DepositFund)
	app.Post("/withdrawFund", accountController.WithdrawFund)
	app.Post("/closeAccount", accountController.CloseAccount)

	app.Listen(":8000")
}

// func main() {

// 	servers := []string{"localhost:9092"}
// 	producer, err := sarama.NewSyncProducer(servers, nil)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer producer.Close()

// 	msg := sarama.ProducerMessage{
// 		Topic: "hello",
// 		Value: sarama.StringEncoder("hello world"),
// 	}

// 	p, o, err := producer.SendMessage(&msg)

// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("patition=%v , offset=%v", p, o)
// }
