package controllers

import (
	"github.com/Acnologla/cdn/internal/core/port"
	"github.com/gin-gonic/gin"
)

type ContentController struct {
	ContentService port.ContentService
}

func (controller *ContentController) Upload(c *gin.Context) {
	// upload file from an urll in req.body.url

}

func NewContentController(contentService port.ContentService) *ContentController {
	return &ContentController{
		ContentService: contentService,
	}
}
