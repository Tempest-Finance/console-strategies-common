package bignumber

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
)

var (
	Zero              = big.NewInt(0)
	One               = big.NewInt(1)
	MaxU128, _        = new(big.Int).SetString("ffffffffffffffffffffffffffffffff", 16)
	ScaleFactorX80    = new(big.Int).Lsh(big.NewInt(1), 64) // 1 << 80
	ScaleFactorX96, _ = new(big.Int).SetString("1000000000000000000000000", 16)
)

// SetFromDecimal256 converts a math.Decimal256 to a big.Int
func SetFromDecimal256(v math.Decimal256) *big.Int {
	bigV := big.Int(v)

	return &bigV
}
