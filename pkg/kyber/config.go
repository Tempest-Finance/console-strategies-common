package kyber

type PriceMapping struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Config struct {
	GetPriceUrl  string
	PriceMapping []PriceMapping
}
