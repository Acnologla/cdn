package port

import (
	"context"
	"io"
)

type Storage interface {
	Upload(ctx context.Context, relativePath string, sekker io.ReadSeeker) error
	Get(ctx context.Context, relativePath string) ([]byte, error)
}
