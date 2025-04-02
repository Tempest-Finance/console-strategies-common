package adapter

import (
	"github.com/Tempest-Finance/console-strategies-common/pkg/ethrpc/adapter/ethereum"
)

func New(url string) (EthClientAdapter, error) {
	return ethereum.NewAdapter(url)
}
