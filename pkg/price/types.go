package price

import (
	"github.com/shopspring/decimal"
)

type Token struct {
	ChainID  int64
	Address  string
	PriceUsd decimal.Decimal
}

type GetTokenPriceRes struct {
	Data map[string]float64 `json:"data"`
}
