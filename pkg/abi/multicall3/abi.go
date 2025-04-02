package multicall3

import "github.com/ethereum/go-ethereum/accounts/abi"

var (
	ABI *abi.ABI
)

func init() {
	ABI, _ = Multicall3MetaData.GetAbi()
}
