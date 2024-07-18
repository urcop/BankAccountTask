package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/urcop/BankAccountTask/domain"
	"github.com/urcop/BankAccountTask/domain/cases/account_create"
	"github.com/urcop/BankAccountTask/domain/cases/account_deposit"
	"github.com/urcop/BankAccountTask/domain/cases/account_get_balance"
	"github.com/urcop/BankAccountTask/domain/cases/account_withdraw"
)

func CreateAccountHandler(c domain.Context, ctx *fiber.Ctx) *RawResponse {
	account, err := account_create.Run(c)
	if err != nil {
		return InternalServerError(err)
	}
	return OK(account)
}

func DepositAccountHandler(c domain.Context, ctx *fiber.Ctx) *RawResponse {
	var request account_deposit.Request
	var param = ctx.Params("id")

	if err := ctx.BodyParser(&request); err != nil {
		return BadRequest(err)
	}

	account, err := account_deposit.Run(c, param, request)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(account)
}

func GetBalanceAccountHandler(c domain.Context, ctx *fiber.Ctx) *RawResponse {
	var param = ctx.Params("id")

	response, err := account_get_balance.Run(c, param)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(response)
}

func WithdrawAccountHandler(c domain.Context, ctx *fiber.Ctx) *RawResponse {
	var request account_withdraw.Request
	var param = ctx.Params("id")

	if err := ctx.BodyParser(&request); err != nil {
		return BadRequest(err)
	}

	account, err := account_withdraw.Run(c, param, request)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(account)
}
