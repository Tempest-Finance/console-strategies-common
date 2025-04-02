package manageroot

import (
	"encoding/hex"

	"github.com/Tempest-Finance/console-strategies-common/pkg/abi/manageroot"
	"github.com/Tempest-Finance/console-strategies-common/pkg/ethrpc"
	"github.com/Tempest-Finance/console-strategies-common/pkg/rpcregistry"
	"github.com/ethereum/go-ethereum/common"
)

type Caller struct {
	rpcRegistry rpcregistry.IRegistry
}

func NewManageRootCaller(rpcRegistry rpcregistry.IRegistry) *Caller {
	return &Caller{rpcRegistry: rpcRegistry}
}

func (c *Caller) GetManageRoot(target string, strategist common.Address, chainId int64) (string, error) {
	rpcClient, err := c.rpcRegistry.GetRpcClient(chainId)
	if err != nil {
		return "", err
	}

	var root [32]byte
	req := rpcClient.NewRequest()
	req.AddCall(&ethrpc.Call{
		ABI:    *manageroot.ABI,
		Target: target,
		Method: "manageRoot",
		Params: []any{strategist},
	}, []any{&root})

	if _, err := req.TryBlockAndAggregate(); err != nil {
		return "", err
	}

	rootHex := "0x" + hex.EncodeToString(root[:])

	return rootHex, nil
}
