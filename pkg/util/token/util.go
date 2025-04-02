package token

import (
	"math"

	"github.com/shopspring/decimal"
)

func MulExpDecimals(value decimal.Decimal, decimals int64) decimal.Decimal {
	return value.Mul(decimal.NewFromFloat(math.Pow(10, float64(decimals))))
}

func DivExpDecimals(value decimal.Decimal, decimals int64) decimal.Decimal {
	return value.Div(decimal.NewFromFloat(math.Pow(10, float64(decimals))))
}

func ExpDecimals(decimals int64) decimal.Decimal {
	return decimal.NewFromFloat(math.Pow(10, float64(decimals)))
}
