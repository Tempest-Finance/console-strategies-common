package rpcregistry

import "github.com/ethereum/go-ethereum/common"

type Config = map[int64]ChainConfig

type ChainConfig struct {
	HTTP             string
	MulticallAddress common.Address
}
