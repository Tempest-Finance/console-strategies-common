package nucleus

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/Tempest-Finance/console-strategies-common/pkg/abi/manageroot"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Tempest-Finance/console-strategies-common/pkg/rpcregistry"
)

type Call struct {
	TargetAddress       string        `json:"target_address"`
	Data                []byte        `json:"data"`
	Value               int           `json:"value"`
	Args                []interface{} `json:"args"`
	FunctionSignature   string        `json:"function_signature"`
	ProofData           []interface{} `json:"proof_data"`
	DecoderAndSanitizer string        `json:"decoder_and_sanitizer"`
}

type CalldataQueue struct {
	client            IClient
	managerAddress    common.Address
	chainId           int64
	strategistAddress string
	root              string
	calls             []Transaction
	rpcRegistry       rpcregistry.IRegistry
	transactor        *bind.TransactOpts
}

func NewCalldataQueue(
	chainId int64,
	strategistAddress string,
	symbol string,
	client IClient,
	rpcRegistry rpcregistry.IRegistry,
	transactor *bind.TransactOpts,
) (*CalldataQueue, error) {
	managerAddress := common.HexToAddress(client.GetAddressBook()[chainId].Nucleus.Vaults[symbol].Manager)
	ethClient, err := rpcRegistry.GetClient(chainId)

	caller, err := manageroot.NewManageRootCaller(managerAddress, ethClient)
	if err != nil {
		return nil, err
	}

	root, err := caller.ManageRoot(nil, common.HexToAddress(strategistAddress))
	if err != nil {
		return nil, err
	}

	rootHex := "0x" + hex.EncodeToString(root[:])

	if !strings.HasPrefix(rootHex, "0x") {
		rootHex = "0x" + rootHex
	}
	zeroRoot := "0x" + strings.Repeat("0", 64)
	if rootHex == zeroRoot {
		return nil, ErrStrategiesIsInvalid
	}

	return &CalldataQueue{
		client:            client,
		managerAddress:    managerAddress,
		chainId:           chainId,
		strategistAddress: strategistAddress,
		root:              rootHex,
		calls:             []Transaction{},
		rpcRegistry:       rpcRegistry,
		transactor:        transactor,
	}, nil
}

func (c *CalldataQueue) AddCall(targetAddress common.Address, calldata []byte, value *big.Int) {
	c.calls = append(c.calls, Transaction{
		Target:    targetAddress,
		Data:      "0x" + common.Bytes2Hex(calldata),
		Val:       value,
		DataBytes: calldata,
	})
}

func (c *CalldataQueue) GetCalldata(ctx context.Context) (*Calldata, error) {
	batchResults, err := c.getBatchProofsAndDecoders(ctx, c.calls)
	if err != nil {
		return nil, err
	}

	var targets []common.Address
	var data [][]byte
	var values []*big.Int

	for _, tx := range c.calls {
		targets = append(targets, tx.Target)
		values = append(values, tx.Val)
		data = append(data, tx.DataBytes)
	}

	return &Calldata{
		ManageProofs:          mappingManageProofs(batchResults.ManageProofs),
		DecodersAndSanitizers: mappingDecodersAndSanitizers(batchResults.DecodersAndSanitizers),
		Targets:               targets,
		TargetData:            data,
		Values:                values,
	}, nil
}

func (c *CalldataQueue) Execute(ctx context.Context) (string, error) {
	client, err := c.rpcRegistry.GetClient(c.chainId)

	if c.calls == nil || len(c.calls) == 0 {
		return "", ErrEmptyCalls
	}

	if !strings.EqualFold(c.transactor.From.Hex(), c.strategistAddress) {
		return "", ErrInvalidSigner
	}

	calldata, err := c.GetCalldata(ctx)
	if err != nil {
		return "", err
	}

	manageRootContract, err := manageroot.NewManageRoot(c.managerAddress, client)
	if err != nil {
		return "", err
	}

	tx, err := manageRootContract.ManageVaultWithMerkleVerification(
		c.transactor,
		calldata.ManageProofs,
		calldata.DecodersAndSanitizers,
		calldata.Targets,
		calldata.TargetData,
		calldata.Values,
	)
	if err != nil {
		return "", err
	}

	txHash, err := c.waitForTransactionSuccess(ctx, tx.Hash(), c.chainId)
	if err != nil {
		return "", err
	}

	return txHash, nil
}

func (c *CalldataQueue) getBatchProofsAndDecoders(ctx context.Context, txs []Transaction) (*MerkleProofs, error) {
	body, err := json.Marshal(map[string]interface{}{
		"chain": strconv.FormatInt(c.chainId, 10),
		"calls": txs,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}
	resp, err := c.client.Post(ctx, MultiproofsUrl+c.root, body)
	if err != nil {
		return nil, fmt.Errorf("failed to get batch proofs and decoders: %w", err)
	}

	var response *MerkleProofs
	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error marshaling input: %v", err)
	}

	if err := json.Unmarshal(jsonBytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response, nil
}

func (c *CalldataQueue) waitForTransactionSuccess(ctx context.Context, txHash common.Hash, chainId int64) (string, error) {
	client, err := c.rpcRegistry.GetClient(chainId)
	if err != nil {
		return "", err
	}
	timeoutTimer := time.NewTimer(60 * time.Second)
	defer timeoutTimer.Stop()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return "", fmt.Errorf("timeout reached while waiting tx")
		case <-ticker.C:
			receipt, err := client.TransactionReceipt(context.Background(), txHash)
			if err != nil {
				log.Printf("getting transaction receipt - error: %v, sleep for 1s...\n", err)
				time.Sleep(time.Second)
			}
			if receipt != nil {
				if receipt.Status == 1 {
					return txHash.String(), nil
				} else {
					return "", fmt.Errorf("transaction failed with status %d", receipt.Status)
				}
			}
		case <-timeoutTimer.C:
			return "", fmt.Errorf("timeout reached while waiting tx")
		}
	}
}
