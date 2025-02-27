package port

import (
	"context"
	"io"

	"github.com/Acnologla/cdn/internal/core/domain"
)

type Storage interface {
	Upload(ctx context.Context, relativePath string, sekker io.ReadSeeker, contentType string) error
	Get(ctx context.Context, relativePath string) (*domain.File, error)
}
