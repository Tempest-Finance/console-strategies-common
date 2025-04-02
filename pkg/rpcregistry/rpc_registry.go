package rpcregistry

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Tempest-Finance/console-strategies-common/pkg/abi/multicall3"
	"github.com/Tempest-Finance/console-strategies-common/pkg/ethrpc"
	"github.com/Tempest-Finance/console-strategies-common/pkg/ethrpc/adapter"
)

// RpcRegistry is a collection of clients for different chains
type RpcRegistry struct {
	clientByChainID    map[int64]*ethclient.Client
	rpcClientByChainID map[int64]*ethrpc.Client
}

func NewRpcRegistry(config Config) (*RpcRegistry, error) {
	clientByChainID := make(map[int64]*ethclient.Client)
	rpcClientByChainID := make(map[int64]*ethrpc.Client)

	for chainID, cfg := range config {
		client, err := ethclient.Dial(cfg.HTTP)
		if err != nil {
			return nil, fmt.Errorf("failed to dial RPC for chainID %d: %w", chainID, err)
		}

		ethClientAdapter, err := adapter.New(cfg.HTTP)
		if err != nil {
			return nil, fmt.Errorf("failed to create eth client adapter for chainID %d: %w", chainID, err)
		}

		rpcClient := ethrpc.NewClient(
			ethrpc.WithEthClientAdapter(ethClientAdapter),
			ethrpc.WithMulticall(cfg.MulticallAddress, multicall3.ABI),
			ethrpc.WithRequestMiddlewares(ethrpc.ParseRequestMiddleware),
			ethrpc.WithResponseMiddlewares(ethrpc.ParseResponseMiddleware),
		)

		clientByChainID[chainID] = client
		rpcClientByChainID[chainID] = rpcClient
	}

	return &RpcRegistry{
		clientByChainID:    clientByChainID,
		rpcClientByChainID: rpcClientByChainID,
	}, nil
}

func (h *RpcRegistry) GetClient(chainID int64) (*ethclient.Client, error) {
	client, ok := h.clientByChainID[chainID]
	if !ok {
		return nil, fmt.Errorf("no client found for chainID %d", chainID)
	}

	return client, nil
}

func (h *RpcRegistry) GetRpcClient(chainID int64) (*ethrpc.Client, error) {
	client, ok := h.rpcClientByChainID[chainID]
	if !ok {
		return nil, fmt.Errorf("no rpc client found for chainID %d", chainID)
	}

	return client, nil
}
