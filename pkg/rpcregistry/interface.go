package rpcregistry

import (
	"github.com/Tempest-Finance/console-strategies-common/pkg/ethrpc"
	"github.com/ethereum/go-ethereum/ethclient"
)

// IRegistry is a collection of clients for different chains
type IRegistry interface {
	// GetClient returns the client for a given chainID
	GetClient(chainID int64) (*ethclient.Client, error)

	// GetRpcClient returns the rpc client for a given chain
	GetRpcClient(chainID int64) (*ethrpc.Client, error)
}
