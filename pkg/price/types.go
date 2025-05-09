package price

type Request struct {
	ChainID      string `json:"chainId"`
	TokenAddress string `json:"tokenAddress"`
	Timestamp    int64  `json:"timestamp"`
}

type Token struct {
	ChainID  string `json:"chainId"`
	Address  string `json:"address"`
	PriceUsd string `json:"priceUsd"`
}

type GetTokenPriceRes struct {
	Data []Token `json:"data"`
}
