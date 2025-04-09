package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/woxQAQ/upload-server/internal/models"
	stores "github.com/woxQAQ/upload-server/internal/stores/progress"
	"github.com/woxQAQ/upload-server/internal/types"
)

type Controller interface {
	Register(g *gin.RouterGroup)
}

type uploadController struct {
	minioClient *minio.Client
	ps          stores.ProgressStore
}

func NewUploadController(minio *minio.Client, ps stores.ProgressStore) Controller {
	return &uploadController{
		minio,
		ps,
	}
}

type PresignRequest struct {
	FileName string `json:"filename"`
	FileSize string `json:"filesize"`
	MD5      string `json:"md5"`
}

func (u *uploadController) Register(e *gin.RouterGroup) {
	e.GET("/presign", u.Presign)
}

func (u *uploadController) Presign(c *gin.Context) {
	ctx := c.Request.Context()
	var req PresignRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	url, err := u.minioClient.PresignedPutObject(
		ctx,
		"file-upload",
		req.FileName,
		10*time.Minute,
	)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	var dtl json.RawMessage

	detail := &types.ImportTaskDetail{
		DatabaseType: types.Mysql,
		Database:     "test",
		Table:        "test",
		OSSUrl:       url,
		MD5:          req.MD5,
		FileName:     req.FileName,
		FileSize:     req.FileSize,
	}

	bytes, err := json.Marshal(detail)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = dtl.UnmarshalJSON(bytes)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	u.ps.CreateProgress(ctx, models.Progress{
		Kind:       types.TaskUpload,
		TaskDetail: dtl,
	})

	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}
