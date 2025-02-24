package main

import (
	"consumer/repositories"
	"consumer/services"
	"context"
	"events"
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/Shopify/sarama.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

func initDatabase() *gorm.DB {

	// 	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// "%v:%v@tcp(%v:%v)/%v",

	dsn := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v",
		viper.GetString("db.username"),
		viper.GetString(("db.password")),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"))
	dial := postgres.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)

	}

	return db
}

func main() {

	// กำหนดการตั้งค่าของ Kafka consumer group
	config := sarama.NewConfig()

	// กำหนดเวอร์ชันของ Kafka ที่ต้องการใช้ (อย่างน้อย V0_10_2_0 ขึ้นไป)
	config.Version = sarama.V0_10_2_0

	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), config)

	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	db := initDatabase()

	accountRepository := repositories.NewAccountRepository(db)
	accountEventHandler := services.NewAccountEventHandler(accountRepository)
	accountConsumerHandler := services.NewConsumerHandler(accountEventHandler)

	fmt.Println("account consumer started...")
	for {
		consumer.Consume(context.Background(), events.Topics, accountConsumerHandler)
	}
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
