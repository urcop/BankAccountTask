package account_create

import (
	"fmt"
	"github.com/urcop/BankAccountTask/domain"
	"github.com/urcop/BankAccountTask/domain/models"
	"math/rand"
)

type Response struct {
	Account *models.Account `json:"account"`
}

func Run(context domain.Context) (*Response, error) {
	account, err := context.Connection().Account().Create(generateId())
	if err != nil {
		return nil, err
	}

	return &Response{Account: account}, nil
}

func generateId() string {
	return fmt.Sprintf("%d", rand.Intn(1000))
}
