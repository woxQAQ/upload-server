package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

type Controller interface {
	Register(g *gin.RouterGroup)
}

type uploadController struct {
	minioClient *minio.Client
}

func NewUploadController(minio *minio.Client) Controller {
	return &uploadController{
		minio,
	}
}

type PresignRequest struct {
	FileName string `json:"filename"`
}

func (u *uploadController) Register(e *gin.RouterGroup) {
	e.GET("/presign", u.Presign)
}

func (u *uploadController) Presign(c *gin.Context) {
	var req *PresignRequest
	err := c.ShouldBind(req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	url, err := u.minioClient.PresignedPutObject(
		c.Request.Context(),
		"file-upload",
		req.FileName,
		10*time.Minute,
	)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}
