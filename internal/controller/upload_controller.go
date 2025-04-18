package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/panjf2000/ants/v2"

	"github.com/woxQAQ/upload-server/internal/models"
	"github.com/woxQAQ/upload-server/internal/stores/progress"
	"github.com/woxQAQ/upload-server/internal/types"
	"github.com/woxQAQ/upload-server/pkg/constants"
)

type uploadController struct {
	mc   *minio.Client
	ps   progress.ProgressStore
	pool *ants.Pool
}

func NewUploadController(minio *minio.Client, ps progress.ProgressStore) Controller {
	pool, err := ants.NewPool(
		1000,
		ants.WithPanicHandler(func(a any) {
			log.Println(a)
		}))
	if err != nil {
		panic(err)
	}
	return &uploadController{
		minio,
		ps,
		pool,
	}
}

func (u *uploadController) Register(e *gin.RouterGroup) {
	e.GET("/presign", u.Presign)
	e.POST("/approve", u.Approve)
	e.POST("/exec", u.Exec)
}

func (u *uploadController) Exec(c *gin.Context) {
	ctx := c.Request.Context()
	var req ExecRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

// Approve godoc
//
//	@Summary		approve a progress
//	@Description	approve progresss
//	@Tags			upload
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Param			req	body	ApproveRequest	false	"request body"
//	@Router			/api/v1/approve [post]
func (u *uploadController) Approve(c *gin.Context) {
	ctx := c.Request.Context()
	var req ApproveRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = u.ps.ApproveProgress(ctx, progress.ApproveOption{
		Id:       req.ProgressId,
		Approver: req.Approver,
		Opinion:  req.Opinion,
	})
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	if req.Trigger == TriggerCurrent {
		execCtx := context.Background()
		u.pool.Submit(func() {
			err = u.Import(execCtx,
				u.mc,
				req.ProgressId,
			)
			if err != nil {
				panic(err)
			}
		})
	}
}

// Presign godoc
//
//	@Summary		presign a data file
//	@Description	presign from oss
//	@Tags			upload
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Param			req	body	PresignRequest	false	"request body"
//	@Router			/api/v1/presign [get]
func (u *uploadController) Presign(c *gin.Context) {
	ctx := c.Request.Context()
	var req PresignRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	url, err := u.mc.PresignedPutObject(
		ctx,
		constants.BucketName,
		req.FileName,
		10*time.Minute,
	)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	var dtl json.RawMessage

	detail := &types.ImportTaskDetail{
		DatabaseType: types.Mysql,
		Database:     req.Database,
		Table:        req.Table,
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
