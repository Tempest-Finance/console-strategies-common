package nucleus

import (
	"encoding/json"
	"strconv"

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

func parseAddressBook(data []byte) (AddressBook, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	book := make(AddressBook)

	for key, val := range raw {
		id, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			continue
		}

		var net NetworkData
		if err := json.Unmarshal(val, &net); err != nil {
			return nil, err
		}
		book[id] = net
	}

	return book, nil
}

func (n *ChainConfig) UnmarshalJSON(data []byte) error {
	var raw rawNucleus
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	n.Vaults = make(map[string]VaultDetail)
	for key, value := range raw {
		if key == "roosterMicroManager" {
			continue
		} else {
			var vault VaultDetail
			if err := json.Unmarshal(value, &vault); err != nil {
				return err
			}
			n.Vaults[key] = vault
		}
	}
	return nil
}
