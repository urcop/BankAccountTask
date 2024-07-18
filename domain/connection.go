package domain

import "github.com/urcop/BankAccountTask/domain/repository/memory"

type Connection interface {
	Account() memory.AccountRepository
}
