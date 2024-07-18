package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	v1 "github.com/urcop/BankAccountTask/api/v1"
	"github.com/urcop/BankAccountTask/domain"
	"github.com/urcop/BankAccountTask/domain/repository"
	"time"
)

var CTX domain.Context

type accountServer struct {
	app *fiber.App
}

func (a accountServer) Start() {
	domainCtx := &ctx{
		connection: repository.Make(),
	}

	a.app.Use(logger.New(logger.Config{
		Format:       "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Europe/Moscow",
		TimeInterval: 500 * time.Millisecond,
	}))

	a.app.Use(recover.New())

	CTX = domainCtx

	a.initializeRoutes(a.app)

	err := a.app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}

func NewAccountServer() Server {
	app := fiber.New()

	return &accountServer{app: app}
}

func (a accountServer) initializeRoutes(app *fiber.App) {
	account := app.Group("/accounts")

	account.Post("/:id/deposit", v1.WrapHandler(v1.DepositAccountHandler, CTX))
	account.Post("/:id/withdraw", v1.WrapHandler(v1.WithdrawAccountHandler, CTX))
	account.Post("/", v1.WrapHandler(v1.CreateAccountHandler, CTX))
	account.Get("/:id/balance", v1.WrapHandler(v1.GetBalanceAccountHandler, CTX))

}
