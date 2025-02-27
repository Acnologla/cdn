package controllers

import (
	"github.com/Acnologla/cdn/internal/core/service"
	"github.com/gin-gonic/gin"
)

type ContentController struct {
	ContentService *service.Content
}

func (controller *ContentController) Get(c *gin.Context) {
	relativePath := c.Param("path")
	if relativePath == "" {
		c.JSON(400, gin.H{"error": "path is required"})
		return
	}
	file, err := controller.ContentService.Get(c.Request.Context(), relativePath)
	if err != nil {
		c.String(404, "File not found")
		return
	}

	c.Data(200, file.ContentType, file.Content)
}

func NewContentController(contentService *service.Content) *ContentController {
	return &ContentController{
		ContentService: contentService,
	}
}
