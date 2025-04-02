package big

import "math/big"

func Max(a *big.Int, b *big.Int) *big.Int {
	if a.Cmp(b) > 0 {
		return a
	}
	return b
}

func Min(a *big.Int, b *big.Int) *big.Int {
	if a.Cmp(b) < 0 {
		return a
	}
	return b
}
