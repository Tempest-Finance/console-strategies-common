package nucleus

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type AddressBook map[int64]NetworkData

type NetworkData struct {
	ID        string            `json:"name"`
	Multisig  string            `json:"multisig,omitempty"`
	Token     map[string]string `json:"token,omitempty"`
	Hyperswap HyperswapData     `json:"hyperswap,omitempty"`
	Nucleus   ChainConfig       `json:"nucleus"`
}

type HyperswapData struct {
	Router                     string `json:"router"`
	NonfungiblePositionManager string `json:"nonfungiblePositionManager"`
}

type VaultDetail struct {
	BoringVault    string `json:"boring_vault"`
	Manager        string `json:"manager"`
	Accountant     string `json:"accountant"`
	Teller         string `json:"teller"`
	RolesAuthority string `json:"roles_authority"`
}

type ChainConfig struct {
	Vaults map[string]VaultDetail `json:"-"`
}

type rawNucleus map[string]json.RawMessage

type Transaction struct {
	Target    common.Address `json:"target"`
	Val       *big.Int       `json:"value"`
	Data      string         `json:"calldata"`
	DataBytes []byte
}

type MerkleProofs struct {
	ManageProofs          [][]string `json:"proofs"`
	DecodersAndSanitizers []string   `json:"decoderAndSanitizerAddress"`
}

type Calldata struct {
	ManageProofs          [][][32]byte     `json:"proofs"`
	DecodersAndSanitizers []common.Address `json:"decoderAndSanitizerAddress"`
	Targets               []common.Address `json:"targets"`
	TargetData            [][]byte         `json:"targetData"`
	Values                []*big.Int       `json:"values"`
}

type StrategistSignerConfig struct {
	Address string `json:"address" mapstructure:"address"`
	Name    string `json:"name" mapstructure:"name"`
}
