package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type FilterQuery struct {
	BlockHash *common.Hash
	FromBlock *big.Int
	ToBlock   *big.Int
	Addresses []common.Address
	Topics    [][]common.Hash
}
