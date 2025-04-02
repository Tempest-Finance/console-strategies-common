package executorplugin

import "github.com/ethereum/go-ethereum/accounts/abi"

var (
	ABI *abi.ABI
)

func init() {
	ABI, _ = ExecutorPluginMetaData.GetAbi()
}
