package httpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/Acnologla/cdn/internal/core/port"
)

type HttpClient struct{}

func (h *HttpClient) Get(ctx context.Context, url string) (io.ReadSeeker, error) {
	response, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	byteSlice := []byte{}
	_, err = response.Body.Read(byteSlice)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(byteSlice), nil
}

func New() port.HttpClient {
	return &HttpClient{}
}
