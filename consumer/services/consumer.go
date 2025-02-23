package services

import "gopkg.in/Shopify/sarama.v1"

type consumerHandler struct {
	// // Setup is run at the beginning of a new session, before ConsumeClaim.
	// Setup(ConsumerGroupSession) error

	// // Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
	// // but before the offsets are committed for the very last time.
	// Cleanup(ConsumerGroupSession) error

	// // ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
	// // Once the Messages() channel is closed, the Handler must finish its processing
	// // loop and exit.
	// ConsumeClaim(ConsumerGroupSession, ConsumerGroupClaim) error

	eventHandler EventHandler
}

func NewConsumerHandler(eventHandler EventHandler) sarama.ConsumerGroupHandler {
	return consumerHandler{eventHandler}
}

func (obj consumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (obj consumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (obj consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for msg := range claim.Messages() {
		obj.eventHandler.Handle(msg.Topic, msg.Value)

		//mark เพื่อบอกว่างานนี้ทำแล้วนะ ไม่งั้น msg จะวนมาให้ทำอีก
		session.MarkMessage(msg, "")
	}

	return nil
}
