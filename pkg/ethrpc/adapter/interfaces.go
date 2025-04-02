package adapter

import (
	"context"
	"github.com/Tempest-Finance/console-strategies-common/pkg/ethrpc/adapter/types"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type EthClientAdapter interface {
	CallContract(ctx context.Context, msg *types.CallMsg, blockNumber *big.Int) ([]byte, error)
	CallContractAtHash(ctx context.Context, msg *types.CallMsg, blockHash common.Hash) ([]byte, error)

	SubscribeNewHead(ctx context.Context, headerChannel chan<- *types.Header) (types.Subscription, error)
	FilterLogs(ctx context.Context, query types.FilterQuery) ([]types.Log, error)

	BlockNumber(ctx context.Context) (uint64, error)
	HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)

	SuggestGasPrice(ctx context.Context) (*big.Int, error)
}
