package brahma

import (
	"os"

	"github.com/Tempest-Finance/console-strategies-common/pkg/crypto"
)

type ExecutorSigner struct {
	signerMap map[string]crypto.ISigner
}

func NewExecutorSigner(config []ExecutorSignerConfig) (*ExecutorSigner, error) {
	signerMap := make(map[string]crypto.ISigner)
	for _, signer := range config {
		s, err := crypto.NewSigner(os.Getenv(signer.Name))
		if err != nil {
			return nil, err
		}
		signerMap[signer.Address] = s
	}
	return &ExecutorSigner{
		signerMap: signerMap,
	}, nil
}

func (e *ExecutorSigner) GetExecutorSigner(address string) (crypto.ISigner, bool) {
	signer, ok := e.signerMap[address]
	return signer, ok
}
