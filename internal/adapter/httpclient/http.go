package httpclient

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/Acnologla/cdn/internal/core/port"
)

type HttpClient struct{}

func (h *HttpClient) IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (h *HttpClient) Get(ctx context.Context, url string) (*http.Response, error) {
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

	return response, nil
}

func NewHttpClient() port.HttpClient {
	return &HttpClient{}
}
