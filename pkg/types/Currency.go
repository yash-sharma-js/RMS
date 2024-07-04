package types

import "time"

type Currency struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Symbol string    `json:"symbol"`
    Base_currency  bool      `json:"base_currency"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
