package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/Acnologla/cdn/internal/core/port"
)

type Content struct {
	storage    port.Storage
	httpClient port.HttpClient
	cdnURL     string
}

func (c *Content) Upload(ctx context.Context, url, path string) (string, error) {
	reader, err := c.httpClient.Get(ctx, url)
	if err != nil {
		return "", err
	}
	err = c.storage.Upload(ctx, path, reader)
	if err != nil {
		return "", err
	}

	fileType := strings.Split(url, ".")[1]
	return fmt.Sprintf("%s/%s.%s", c.cdnURL, path, fileType), nil
}
