package repository

import "github.com/urcop/BankAccountTask/domain/repository/memory"

type Connection struct {
	accountRepository memory.AccountRepository
}

func (c Connection) Account() memory.AccountRepository {
	return c.accountRepository
}

func Make() *Connection {
	return &Connection{
		accountRepository: memory.NewAccountRepository(),
	}
}
