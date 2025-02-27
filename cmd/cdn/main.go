package main

import (
	"context"

	"github.com/Acnologla/cdn/internal/adapter/config"
	"github.com/Acnologla/cdn/internal/adapter/http"
	"github.com/Acnologla/cdn/internal/adapter/http/controllers"
	"github.com/Acnologla/cdn/internal/adapter/httpclient"
	"github.com/Acnologla/cdn/internal/adapter/storage"
	"github.com/Acnologla/cdn/internal/core/service"
)

func main() {
	config, err := config.LoadConfig()
	context := context.Background()
	if err != nil {
		panic(err)
	}

	// initialize adapters

	storageAdapter := storage.NewWasabiStorage(context, config.WasabiConfig)
	//cacheAdapter := cache.NewLRUCache(512)
	httpClientAdapter := httpclient.NewHttpClient()

	// initialize services

	contentService := service.NewContentService(storageAdapter, httpClientAdapter, config.HTTPConfig.ServerURL)

	//initialize controllers

	contentController := controllers.NewContentController(contentService)

	// initialize http server

	err = http.CreateAndServe(config.HTTPConfig, contentController)
	if err != nil {
		panic(err)
	}
}
