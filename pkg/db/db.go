package db

import (
	"database/sql"

	"github.com/currency/pkg/types"
)

type PostgresCurrencyRepository struct {
	DB *sql.DB
}

func NewPostgresCurrencyRepository(db *sql.DB) *PostgresCurrencyRepository {
	return &PostgresCurrencyRepository{DB: db}
}

func (r *PostgresCurrencyRepository) Create(currency *types.Currency) error {
	query := `INSERT INTO master_currency (name, symbol, base_currency, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return r.DB.QueryRow(query, currency.Name, currency.Symbol, currency.Base_currency, currency.CreatedAt, currency.UpdatedAt).Scan(&currency.ID)
}

func (r *PostgresCurrencyRepository) GetAll(id *int, name *string) ([]*types.Currency, error) {
	var query string
	var args []interface{}

	if id != nil {
		query = `SELECT * FROM master_currency WHERE id = $1`
		args = append(args, *id)
	} else if name != nil {
		query = `SELECT * FROM master_currency WHERE name = $1`
		args = append(args, *name)
	} else {
		query = `SELECT * FROM master_currency`
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	currencys := []*types.Currency{}
	for rows.Next() {
		currency := new(types.Currency)
		err := rows.Scan(&currency.ID, &currency.Name, &currency.Symbol, &currency.Base_currency, &currency.CreatedAt, &currency.UpdatedAt)
		if err != nil {
			return nil, err
		}
		currencys = append(currencys, currency)
	}
	return currencys, nil
}


func (r *PostgresCurrencyRepository) GetByID(id int) (*types.Currency, error) {
	query := `SELECT * FROM master_currency WHERE id = $1`
	currency := new(types.Currency)
	err := r.DB.QueryRow(query, id).Scan(&currency.ID, &currency.Name, &currency.Symbol, &currency.Base_currency, &currency.CreatedAt, &currency.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return currency, nil
}

func (r *PostgresCurrencyRepository) GetByNAME(name string) (*types.Currency, error) {
	query := `SELECT * FROM master_currency WHERE name = $1`
	currency := new(types.Currency)
	err := r.DB.QueryRow(query, name).Scan(&currency.ID, &currency.Name, &currency.Symbol, &currency.Base_currency, &currency.CreatedAt, &currency.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return currency, nil
}

func (r *PostgresCurrencyRepository) Update(currency *types.Currency) error {
	query := `UPDATE master_currency SET name = $1, symbol = $2, base_currency = $3, updated_at = $4 WHERE id = $5`
	_, err := r.DB.Exec(query, currency.Name, currency.Symbol, currency.Base_currency, currency.UpdatedAt, currency.ID)
	return err
}

func (r *PostgresCurrencyRepository) Delete(id int) error {
	query := `DELETE FROM master_currency WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
