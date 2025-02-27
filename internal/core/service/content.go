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
	split := strings.Split(url, ".")
	fileType := split[len(split)-1]
	filePath := fmt.Sprintf("%s.%s", path, fileType)
	err = c.storage.Upload(ctx, filePath, reader)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", c.cdnURL, filePath), nil
}

func NewContentService(storage port.Storage, httpClient port.HttpClient, cdnURL string) *Content {
	return &Content{
		storage:    storage,
		httpClient: httpClient,
		cdnURL:     cdnURL,
	}
}
