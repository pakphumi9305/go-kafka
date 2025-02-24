package services

import (
	"errors"
	"events"
	"log"
	"producer/commands"

	"github.com/google/uuid"
)

type AccountService interface {
	OpenAccount(command commands.OpenAccontCommand) (id string, err error)
	DepositeFund(command commands.DepositFundCommand) error
	WithdrawFund(command commands.WithdrawFundCommand) error
	CloseAccount(command commands.CloseAccountCommand) error
}

type accountService struct {
	eventProducer EventProducer
}

func NewAccountService(eventProducer EventProducer) AccountService {
	return accountService{eventProducer}
}

// CloseAccount implements AccountService.
func (a accountService) CloseAccount(command commands.CloseAccountCommand) error {
	//panic("unimplemented")
	if command.ID == "" {
		return errors.New("bad request")
	}

	event := events.CloseAccountEvent{
		ID: command.ID,
	}
	log.Printf("%#v", event)
	err := a.eventProducer.Produce(event)

	return err
}

// DepositeFund implements AccountService.
func (a accountService) DepositeFund(command commands.DepositFundCommand) error {
	//panic("unimplemented")
	if command.ID == "" || command.Amount == 0 {
		return errors.New("bad request")
	}
	event := events.DepositFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}
	log.Printf("%#v", event)

	err := a.eventProducer.Produce(event)
	return err
}

// OpenAccount implements AccountService.
func (a accountService) OpenAccount(command commands.OpenAccontCommand) (id string, err error) {

	if command.AccountHolder == "" || command.AccountType == 0 || command.OpeningBalance == 0 {
		return "", errors.New("bad request")
	}

	event := events.OpenAccontEvent{
		ID:             uuid.NewString(),
		AccountHolder:  command.AccountHolder,
		AccountType:    command.AccountType,
		OpeningBalance: command.OpeningBalance,
	}
	log.Printf("%#v", event)

	err = a.eventProducer.Produce(event)

	return event.ID, err
}

// WithdrawFund implements AccountService.
func (a accountService) WithdrawFund(command commands.WithdrawFundCommand) error {
	//panic("unimplemented")

	if command.ID == "" || command.Amount == 0 {
		return errors.New("bad request")
	}
	event := events.WithdrawFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}

	log.Printf("%#v", event)

	err := a.eventProducer.Produce(event)

	return err
}
