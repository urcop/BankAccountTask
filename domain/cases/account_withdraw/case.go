package account_withdraw

import (
	"fmt"
	"github.com/urcop/BankAccountTask/domain"
	"github.com/urcop/BankAccountTask/domain/models"
	"log"
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

	if account.Balance < r.Amount {
		log.Printf("Account %s: Insufficient funds for withdrawal of %.2f", account.ID, r.Amount)
		return nil, fmt.Errorf("insufficient funds")
	}

	account.Balance -= r.Amount

	if err := context.Connection().Account().Update(accountId, account); err != nil {
		return nil, err
	}

	return &Response{Account: account}, nil
}
