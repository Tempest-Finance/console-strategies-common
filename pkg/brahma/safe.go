package brahma

import (
	"github.com/Brahma-fi/go-safe/encoders"
	"github.com/Brahma-fi/go-safe/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
)

func GetEncodedSafeTx(
	safeMultiSendAddress common.Address,
	safeMultiSendAbi *abi.ABI,
	transactions []types.Transaction,
) (*types.SafeTx, error) {
	packedTransactions, value, err := encoders.PackTransactions(
		&types.SafeMultiSendRequest{
			Transactions: transactions,
		},
	)
	if err != nil {
		return nil, err
	}
	callData, err := encoders.GetEncodedMultiSendTransaction(packedTransactions, safeMultiSendAbi)
	if err != nil {
		return nil, err
	}
	return &types.SafeTx{
		Operation: uint8(1),
		To:        common.NewMixedcaseAddress(safeMultiSendAddress),
		Value:     math.Decimal256(*value),
		Data:      (*hexutil.Bytes)(&callData),
	}, nil
}
