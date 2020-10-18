package transaction

import "time"

type CreateTransaction struct {
	Amount float64 `json:"amount,omitempty"`
	Type   string  `json:"type,omitempty"`
}

type Transaction struct {
	ID            string    `json:"id,omitempty"`
	Amount        float64   `json:"amount,omitempty"`
	Type          string    `json:"type,omitempty"`
	EffectiveDate time.Time `json:"effectiveDate,omitempty"`
}
