package ethrpc

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type RequestExecutor interface {
	Execute(request *Request) (*Response, error)
	GetMulticallContractAddress() common.Address
	GetMulticallABI() *abi.ABI
}
