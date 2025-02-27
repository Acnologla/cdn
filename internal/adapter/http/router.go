package http

import (
	"fmt"

	"github.com/Acnologla/cdn/internal/adapter/config"
	"github.com/Acnologla/cdn/internal/adapter/http/controllers"
	"github.com/Acnologla/cdn/internal/adapter/http/middlewares"
	"github.com/gin-gonic/gin"
)

func CreateAndServe(c config.HTTPConfig, content *controllers.ContentController, contentManagement *controllers.ContentManagementController) error {
	r := gin.New()

	//private routes for uploading and deleting files
	contentManagementControllerGroup := r.Group("/api")
	contentManagementControllerGroup.Use(middlewares.IsAdminMiddleware(c.AdminKey))
	contentManagementControllerGroup.POST("/upload", contentManagement.Upload)

	//public routes for getting files
	r.GET("/cdn/*path", content.Get)

	return r.Run(fmt.Sprintf(":%s", c.Port))
}
