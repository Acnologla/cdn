package httpclient

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/Acnologla/cdn/internal/core/port"
)

type HttpClient struct{}

func (h *HttpClient) IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (h *HttpClient) Get(ctx context.Context, url string) (io.ReadSeeker, error) {
	if !h.IsUrl(url) {
		return nil, errors.New("invalid url")
	}
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	byteSlice, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(byteSlice), nil
}

func NewHttpClient() port.HttpClient {
	return &HttpClient{}
}
