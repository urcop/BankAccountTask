package server

import "github.com/urcop/BankAccountTask/domain"

type ctx struct {
	connection domain.Connection
}

func (c *ctx) Connection() domain.Connection {
	return c.connection
}

func (c *ctx) MakeContext() domain.Context {
	return &ctx{
		connection: c.connection,
	}
}
