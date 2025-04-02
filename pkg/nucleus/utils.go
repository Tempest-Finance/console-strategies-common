package nucleus

import (
	"github.com/ethereum/go-ethereum/common"
)

func mappingDecodersAndSanitizers(proofs []string) []common.Address {
	addresses := make([]common.Address, len(proofs))
	for i, proof := range proofs {
		addresses[i] = common.HexToAddress(proof)
	}
	return addresses
}

func mappingManageProofs(proofs [][]string) [][][32]byte {
	manageProofs := make([][][32]byte, len(proofs))
	for i, proof := range proofs {
		manageProofs[i] = make([][32]byte, len(proof))
		for j, p := range proof {
			manageProofs[i][j] = common.HexToHash(p)
		}
	}
	return manageProofs
}
