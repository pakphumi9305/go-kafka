package controllers

import (
	"log"
	"producer/commands"
	"producer/services"

	"github.com/gofiber/fiber/v2"
)

type AccountController interface {
	OpenAccount(c *fiber.Ctx) error
	DepositFund(c *fiber.Ctx) error
	WithdrawFund(c *fiber.Ctx) error
	CloseAccount(c *fiber.Ctx) error
}

type accountController struct {
	accountService services.AccountService
}

func NewAccountController(accountService services.AccountService) AccountController {
	return accountController{accountService}
}

func (a accountController) OpenAccount(c *fiber.Ctx) error {
	command := commands.OpenAccontCommand{}

	err := c.BodyParser(&command)

	if err != nil {
		return err
	}

	id, err := a.accountService.OpenAccount(command)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusCreated)

	return c.JSON(fiber.Map{
		"message": "open account successfully",
		"id":      id,
	})
}

func (a accountController) DepositFund(c *fiber.Ctx) error {

	command := commands.DepositFundCommand{}

	err := c.BodyParser(&command)

	if err != nil {
		return err
	}

	err = a.accountService.DepositeFund(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"message": "deposit account successfully",
	})

}

func (a accountController) WithdrawFund(c *fiber.Ctx) error {

	command := commands.WithdrawFundCommand{}

	err := c.BodyParser(&command)

	if err != nil {
		return err
	}

	err = a.accountService.WithdrawFund(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"message": "withdraw account successfully",
	})

}

func (a accountController) CloseAccount(c *fiber.Ctx) error {

	command := commands.CloseAccountCommand{}

	err := c.BodyParser(&command)

	if err != nil {
		return err
	}

	err = a.accountService.CloseAccount(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"message": "close account account successfully",
	})

}
