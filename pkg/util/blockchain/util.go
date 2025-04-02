package blockchain

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func IsAddress(s string) bool {
	return common.IsHexAddress(Add0xPrefix(s))
}

func NormalizeAddress(address string) string {
	return strings.ToLower(address)
}

func Add0xPrefix(str string) string {
	if strings.HasPrefix(str, "0x") {
		return str
	}
	return "0x" + str
}

func RemoveOxPrefix(str string) string {
	if !strings.HasPrefix(str, "0x") {
		return str
	}
	return str[2:]
}

func PaddingHex(s string) string {
	// "0x..." -> "0x0..." fill leading 0 with 64 characters
	if len(s) == 66 {
		return s
	}
	return fmt.Sprintf("%s%s%s", "0x", strings.Repeat("0", 66-len(s)), s[2:])
}

func ConcatChainIDAddress(chainID int64, address string) string {
	return fmt.Sprintf("%d:%s", chainID, address)
}
