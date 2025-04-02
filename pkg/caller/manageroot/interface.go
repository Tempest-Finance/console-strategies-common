package manageroot

import (
	"github.com/ethereum/go-ethereum/common"
)

type ICaller interface {
	GetManageRoot(target string, strategist common.Address, chainId int64) (string, error)
}
