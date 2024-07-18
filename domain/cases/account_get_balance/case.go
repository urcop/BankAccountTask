package account_get_balance

import (
	"github.com/urcop/BankAccountTask/domain"
	"github.com/urcop/BankAccountTask/domain/models"
	"sync"
)

var mutex sync.Mutex

type Response struct {
	Account *models.Account `json:"account"`
}

func Run(context domain.Context, accountId string) (*Response, error) {
	mutex.Lock()
	defer mutex.Unlock()

	account, err := context.Connection().Account().GetById(accountId)
	if err != nil {
		return nil, err
	}

	return &Response{Account: account}, nil
}
