package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	adaptertypes "github.com/Tempest-Finance/console-strategies-common/pkg/ethrpc/adapter/types"
)

type Adapter struct {
	client *ethclient.Client
}

func NewAdapter(url string) (*Adapter, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		client: client,
	}, nil
}

func (a *Adapter) CallContract(ctx context.Context, msg *adaptertypes.CallMsg, blockNumber *big.Int) ([]byte, error) {
	ethereumCallMsg := a.convertToEthereumCallMsg(msg)

	return a.client.CallContract(ctx, ethereumCallMsg, blockNumber)
}

func (a *Adapter) CallContractAtHash(ctx context.Context, msg *adaptertypes.CallMsg, blockHash common.Hash) ([]byte, error) {
	ethereumCallMsg := a.convertToEthereumCallMsg(msg)

	return a.client.CallContractAtHash(ctx, ethereumCallMsg, blockHash)
}

func (a *Adapter) SubscribeNewHead(ctx context.Context, headerChannel chan<- *adaptertypes.Header) (adaptertypes.Subscription, error) {
	originHeaderChannel := make(chan *types.Header)
	sub, err := a.client.SubscribeNewHead(ctx, originHeaderChannel)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(originHeaderChannel)

		for {
			select {
			case <-ctx.Done():
				return
			case originHeader := <-originHeaderChannel:
				headerChannel <- a.convertFromEthereumHeader(originHeader)
			}
		}
	}()

	return sub, nil
}

func (a *Adapter) FilterLogs(ctx context.Context, query adaptertypes.FilterQuery) ([]adaptertypes.Log, error) {
	logs, err := a.client.FilterLogs(ctx, a.convertToEthereumFilterQuery(query))
	if err != nil {
		return nil, err
	}

	return a.convertFromEthereumLogs(logs), nil
}

func (a *Adapter) BlockNumber(ctx context.Context) (uint64, error) {
	return a.client.BlockNumber(ctx)
}

func (a *Adapter) HeaderByHash(ctx context.Context, hash common.Hash) (*adaptertypes.Header, error) {
	originHeader, err := a.client.HeaderByHash(ctx, hash)
	if err != nil {
		return nil, err
	}

	return a.convertFromEthereumHeader(originHeader), nil
}

func (a *Adapter) HeaderByNumber(ctx context.Context, number *big.Int) (*adaptertypes.Header, error) {
	originHeader, err := a.client.HeaderByNumber(ctx, number)
	if err != nil {
		return nil, err
	}

	return a.convertFromEthereumHeader(originHeader), nil
}

func (a *Adapter) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return a.client.SuggestGasPrice(ctx)
}

func (a *Adapter) convertToEthereumCallMsg(originMsg *adaptertypes.CallMsg) ethereum.CallMsg {
	return ethereum.CallMsg{
		From:       originMsg.From,
		To:         originMsg.To,
		Gas:        originMsg.Gas,
		GasPrice:   originMsg.GasPrice,
		GasFeeCap:  originMsg.GasFeeCap,
		GasTipCap:  originMsg.GasTipCap,
		Value:      originMsg.Value,
		Data:       originMsg.Data,
		AccessList: a.convertToEthereumAccessList(originMsg.AccessList),
	}
}

func (_ *Adapter) convertToEthereumAccessList(originAccessList adaptertypes.AccessList) types.AccessList {
	accessList := make([]types.AccessTuple, 0, len(originAccessList))

	for _, originAccessTuple := range originAccessList {
		accessList = append(accessList, types.AccessTuple{
			Address:     originAccessTuple.Address,
			StorageKeys: originAccessTuple.StorageKeys,
		})
	}

	return accessList
}

func (_ *Adapter) convertToEthereumFilterQuery(originFilterQuery adaptertypes.FilterQuery) ethereum.FilterQuery {
	return ethereum.FilterQuery{
		BlockHash: originFilterQuery.BlockHash,
		FromBlock: originFilterQuery.FromBlock,
		ToBlock:   originFilterQuery.ToBlock,
		Addresses: originFilterQuery.Addresses,
		Topics:    originFilterQuery.Topics,
	}
}

func (_ *Adapter) convertFromEthereumHeader(originHeader *types.Header) *adaptertypes.Header {
	return &adaptertypes.Header{
		Hash:       originHeader.Hash(),
		ParentHash: originHeader.ParentHash,
		Number:     originHeader.Number,
		Time:       originHeader.Time,
	}
}

func (_ *Adapter) convertFromEthereumLogs(originLogs []types.Log) []adaptertypes.Log {
	logs := make([]adaptertypes.Log, 0, len(originLogs))
	for i := range originLogs {
		topics := make([]common.Hash, 0, len(originLogs[i].Topics))

		topics = append(topics, originLogs[i].Topics...)

		logs = append(logs, adaptertypes.Log{
			Address:     originLogs[i].Address,
			Topics:      topics,
			Data:        originLogs[i].Data,
			BlockNumber: originLogs[i].BlockNumber,
			TxHash:      originLogs[i].TxHash,
			TxIndex:     originLogs[i].TxIndex,
			BlockHash:   originLogs[i].BlockHash,
			Index:       originLogs[i].Index,
			Removed:     originLogs[i].Removed,
		})
	}

	return logs
}
