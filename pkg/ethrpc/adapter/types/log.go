package types

import (
	"github.com/ethereum/go-ethereum/common"
)

type Log struct {
	Address     common.Address `json:"address"`
	Topics      []common.Hash  `json:"topics"`
	Data        []byte         `json:"data"`
	BlockNumber uint64         `json:"blockNumber"`
	TxHash      common.Hash    `json:"transactionHash"`
	TxIndex     uint           `json:"transactionIndex"`
	BlockHash   common.Hash    `json:"blockHash"`
	Index       uint           `json:"logIndex"`
	Removed     bool           `json:"removed"`
}
