package controllers

import (
	"github.com/Acnologla/cdn/internal/core/service"
	"github.com/gin-gonic/gin"
)

type UploadRequest struct {
	URL  string `json:"url" binding:"required"`
	Path string `json:"path" binding:"required"`
}

type ContentManagementController struct {
	ContentService *service.Content
}

func (controller *ContentManagementController) Upload(c *gin.Context) {
	var request UploadRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resultURL, err := controller.ContentService.Upload(c.Request.Context(), request.URL, request.Path)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"url": resultURL})
}

func NewContentManagementController(contentService *service.Content) *ContentManagementController {
	return &ContentManagementController{
		ContentService: contentService,
	}
}
