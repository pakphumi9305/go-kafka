package services

import (
	"consumer/repositories"
	"encoding/json"
	"events"
	"log"
	"reflect"
)

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}

type accountEventHandler struct {
	accountRepo repositories.AccountRepository
}

func NewAccountEventHandler(accountRepo repositories.AccountRepository) EventHandler {
	return accountEventHandler{accountRepo}
}

func (obj accountEventHandler) Handle(topic string, eventBytes []byte) {
	switch topic {
	case reflect.TypeOf(events.OpenAccontEvent{}).Name():
		event := &events.OpenAccontEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount := repositories.BankAccount{
			ID:            event.ID,
			AccountHolder: event.AccountHolder,
			AccountType:   event.AccountType,
			Balance:       event.OpeningBalance,
		}

		err = obj.accountRepo.Save(bankAccount)
		if err != nil {
			log.Println(err)
			return
		}

	case reflect.TypeOf(events.DepositFundEvent{}).Name():
		event := &events.DepositFundEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount, err := obj.accountRepo.FindByID(event.ID)

		if err != nil {
			log.Println(err)
			return
		}
		bankAccount.Balance += event.Amount

		err = obj.accountRepo.Save(bankAccount)
		if err != nil {
			log.Println(err)
			return
		}

	case reflect.TypeOf(events.WithdrawFundEvent{}).Name():

		event := &events.WithdrawFundEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount, err := obj.accountRepo.FindByID(event.ID)

		if err != nil {
			log.Println(err)
			return
		}
		bankAccount.Balance -= event.Amount
		if bankAccount.Balance < 0 {
			log.Println("withdraw exceed your balance")
			return
		}
		err = obj.accountRepo.Save(bankAccount)
		if err != nil {
			log.Println(err)
			return
		}
	case reflect.TypeOf(events.CloseAccountEvent{}).Name():

		event := &events.CloseAccountEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}

		err = obj.accountRepo.Delete(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
	default:
		log.Println("no event handler")
	}
}
