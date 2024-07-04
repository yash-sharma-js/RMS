package service

import (
	"github.com/currency/pkg/db"
	"github.com/currency/pkg/types"
)

type CurrencyService interface {
	CreateCurrency(currency *types.Currency) error
	GetCurrencys(id *int, name *string) ([]*types.Currency, error)
	// GetCurrencyByID(id int) (*types.Currency, error)
	// GetCurrencyByNAME(name string) (*types.Currency, error)
	UpdateCurrency(currency *types.Currency) error
	DeleteCurrency(id int) error
}

type currencyService struct {
	repo db.CurrencyRepository
}

func NewCurrencyService(repo db.CurrencyRepository) CurrencyService {
	return &currencyService{repo: repo}
}

func (s *currencyService) CreateCurrency(currency *types.Currency) error {
	return s.repo.Create(currency)
}

func (s *currencyService) GetCurrencys(id *int, name *string) ([]*types.Currency, error) {
	return s.repo.GetAll(id,name)
}

// func (s *currencyService) GetCurrencyByID(id int) (*types.Currency, error) {
// 	return s.repo.GetByID(id)
// }

// func (s *currencyService) GetCurrencyByNAME(name string) (*types.Currency, error) {
// 	return s.repo.GetByNAME(name)
// }

func (s *currencyService) UpdateCurrency(currency *types.Currency) error {
	return s.repo.Update(currency)
}

func (s *currencyService) DeleteCurrency(id int) error {
	return s.repo.Delete(id)
}
