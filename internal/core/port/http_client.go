package port

import (
	"context"
	"io"
)

type HttpClient interface {
	Get(ctx context.Context, url string) (io.ReadSeeker, error)
}
