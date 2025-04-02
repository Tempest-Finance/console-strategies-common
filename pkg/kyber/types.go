package kyber

import (
	"github.com/Tempest-Finance/console-strategies-common/pkg/util/blockchain"
	"github.com/shopspring/decimal"
)

type TokenConversion struct {
	ChainID int64
	Address string
}

type Token struct {
	ChainID  int64
	Address  string
	PriceUsd decimal.Decimal
}

type GetTokenPriceDataResult struct {
	PriceBuy  decimal.Decimal `json:"PriceBuy"`
	PriceSell decimal.Decimal `json:"PriceSell"`
}

func (r *GetTokenPriceDataResult) ToValueObject(chainID int64, address string) Token {
	price := r.PriceBuy
	if price.IsZero() {
		price = r.PriceSell
	}
	return Token{
		Address:  blockchain.NormalizeAddress(address),
		ChainID:  chainID,
		PriceUsd: price,
	}
}

type GetTokenPriceDataAddress = map[string]GetTokenPriceDataResult

type GetTokenPriceDataChain = map[int64]GetTokenPriceDataAddress

type GetTokenPriceRes struct {
	Data GetTokenPriceDataChain `json:"data"`
}

func (r *GetTokenPriceRes) ToValueObjects() []Token {
	result := make([]Token, 0)
	for chainIDStr, chainData := range r.Data {
		chainID := chainIDStr
		for address, tokenData := range chainData {
			result = append(result, tokenData.ToValueObject(chainID, address))
		}
	}
	return result
}
