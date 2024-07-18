package memory

import (
	"errors"
	"github.com/urcop/BankAccountTask/domain/models"
)

type AccountRepository interface {
	Create(id string) (*models.Account, error)
	GetById(id string) (*models.Account, error)
	Update(id string, account *models.Account) error
}

type AccountRepositoryImpl struct {
	accounts map[string]*models.Account
}

func (a AccountRepositoryImpl) Update(id string, account *models.Account) error {
	existing, err := a.GetById(id)
	if err != nil {
		return err
	}

	a.accounts[id] = &models.Account{
		ID:      existing.ID,
		Balance: account.Balance,
	}
	return nil
}

func (a AccountRepositoryImpl) Create(id string) (*models.Account, error) {
	account := &models.Account{
		ID:      id,
		Balance: 0,
	}
	a.accounts[id] = account
	return account, nil
}

func (a AccountRepositoryImpl) GetById(id string) (*models.Account, error) {
	account, exists := a.accounts[id]
	if !exists {
		return nil, errors.New("account not found")
	}
	return account, nil
}

func NewAccountRepository() AccountRepository {
	return &AccountRepositoryImpl{
		accounts: make(map[string]*models.Account),
	}
}
