package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"path"

	"github.com/Acnologla/cdn/internal/core/domain"
	"github.com/Acnologla/cdn/internal/core/port"
)

type Content struct {
	storage    port.Storage
	cache      port.Cache
	httpClient port.HttpClient
	cdnURL     string
}

func (c *Content) getUploadBody(ctx context.Context, url string) (io.ReadSeeker, string, error) {
	response, err := c.httpClient.Get(ctx, url)
	if err != nil {
		return nil, "", err
	}
	defer response.Body.Close()
	byteSlice, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, "", err
	}
	reader := bytes.NewReader(byteSlice)
	return reader, response.Header.Get("Content-Type"), nil
}

func (c *Content) Upload(ctx context.Context, url, relativePath string) (string, error) {
	reader, contentType, err := c.getUploadBody(ctx, url)
	if err != nil {
		return "", err
	}

	fileType := path.Ext(url)
	filePath := relativePath + fileType
	err = c.storage.Upload(ctx, filePath, reader, contentType)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", c.cdnURL, filePath), nil
}

func (c *Content) Get(ctx context.Context, path string) (*domain.File, error) {
	cacheValue, ok := c.cache.Get(path)
	if ok {
		return cacheValue, nil
	}

	image, err := c.storage.Get(ctx, path)
	if err != nil {
		return nil, err
	}

	c.cache.Set(path, image)
	return image, nil
}

func NewContentService(storage port.Storage, httpClient port.HttpClient, cache port.Cache, cdnURL string) *Content {
	return &Content{
		storage:    storage,
		httpClient: httpClient,
		cdnURL:     cdnURL,
		cache:      cache,
	}
}
