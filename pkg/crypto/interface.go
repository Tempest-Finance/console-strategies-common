package crypto

import (
	"context"
)

type ISigner interface {
	Sign(ctx context.Context, hash []byte) ([]byte, error)
	PublicKey() string
}
