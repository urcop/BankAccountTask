package account_deposit

import (
	"github.com/urcop/BankAccountTask/domain"
	"github.com/urcop/BankAccountTask/domain/models"
	"sync"
)

var mutex sync.Mutex

type Request struct {
	Amount float64 `json:"amount"`
}

type Response struct {
	Account *models.Account `json:"account"`
}

func Run(context domain.Context, accountId string, r Request) (*Response, error) {
	mutex.Lock()
	defer mutex.Unlock()

	account, err := context.Connection().Account().GetById(accountId)
	if err != nil {
		return nil, err
	}

	account.Balance += r.Amount

	if err := context.Connection().Account().Update(accountId, account); err != nil {
		return nil, err
	}

	return &Response{Account: account}, nil
}
