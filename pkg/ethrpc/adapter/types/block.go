package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Block struct {
	Number      *big.Int    `json:"number"`
	Hash        common.Hash `json:"hash"`
	Timestamp   uint64      `json:"timestamp"`
	ParentHash  common.Hash `json:"parentHash"`
	ReorgedHash common.Hash `json:"reorgedHash"`
	Logs        []Log       `json:"logs"`
}
