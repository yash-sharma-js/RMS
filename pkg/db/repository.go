package db

import "github.com/currency/pkg/types"

type CurrencyRepository interface {
    Create(currency *types.Currency) error
    GetAll(id *int, name *string) ([]*types.Currency, error)
    // GetByID(id int) ([]*types.Currency, error)
    // GetByNAME(name string) (*types.Currency, error)
    Update(currency *types.Currency) error
    Delete(id int) error
}
