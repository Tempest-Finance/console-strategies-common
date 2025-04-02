package nucleus

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type NetworkData struct {
	ID          string               `json:"id"`
	Multisig    string               `json:"multisig"`
	Token       map[string]string    `json:"token"`
	Thunderhead ThunderheadData      `json:"thunderhead"`
	Hyperswap   HyperswapData        `json:"hyperswap"`
	Nucleus     map[string]VaultData `json:"nucleus"`
}

type ThunderheadData struct {
	Overseer string `json:"overseer"`
}

type HyperswapData struct {
	Router                     string `json:"router"`
	NonfungiblePositionManager string `json:"nonfungiblePositionManager"`
}

type VaultData struct {
	BoringVault    string `json:"boring_vault"`
	Manager        string `json:"manager"`
	Accountant     string `json:"accountant"`
	Teller         string `json:"teller"`
	RolesAuthority string `json:"roles_authority"`
}

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
