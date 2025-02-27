package controllers

import (
	"github.com/Acnologla/cdn/internal/core/service"
	"github.com/gin-gonic/gin"
)

type UploadRequest struct {
	URL  string `json:"url"`
	Path string `json:"path"`
}

type ContentController struct {
	ContentService service.Content
}

func (controller *ContentController) Upload(c *gin.Context) {
	var request UploadRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	controller.ContentService.Upload(c.Request.Context(), request.URL, request.Path)

	// return the result url
	c.JSON(200, gin.H{"message": "File uploaded"})
}

func NewContentController(contentService service.Content) *ContentController {
	return &ContentController{
		ContentService: contentService,
	}
}
