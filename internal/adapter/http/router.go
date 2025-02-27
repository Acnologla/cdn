package http

import (
	"fmt"

	"github.com/Acnologla/cdn/internal/adapter/config"
	"github.com/Acnologla/cdn/internal/adapter/http/controllers"
	"github.com/Acnologla/cdn/internal/adapter/http/middlewares"
	"github.com/gin-gonic/gin"
)

func CreateAndServe(c config.HTTPConfig, content *controllers.ContentController) error {
	r := gin.New()
	contentController := r.Group("/api")
	contentController.Use(middlewares.IsAdminMiddleware(c.AdminKey))
	contentController.POST("/upload", content.Upload)

	return r.Run(fmt.Sprintf(":%s", c.Port))
}
