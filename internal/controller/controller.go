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
	stores "github.com/woxQAQ/upload-server/internal/stores/progress"
	"github.com/woxQAQ/upload-server/internal/types"
	"github.com/woxQAQ/upload-server/pkg/constants"
)

type Controller interface {
	Register(g *gin.RouterGroup)
}

type uploadController struct {
	mc   *minio.Client
	ps   stores.ProgressStore
	pool *ants.Pool
}

func NewUploadController(minio *minio.Client, ps stores.ProgressStore) Controller {
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
	e.POST("/start", u.Start)
}

func (u *uploadController) Start(c *gin.Context) {

}

// Approve godoc
// @Summary approve a progress
// @Description approve progresss
// @Tags upload
// @Accept json
// @Produce json
// @Success 200
// @Param req body ApproveRequest false "request body"
// @Router /api/v1/approve [post]
func (u *uploadController) Approve(c *gin.Context) {
	ctx := c.Request.Context()
	var req ApproveRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = u.ps.ApproveProgress(ctx, stores.ApproveOption{
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
			err = Import(execCtx,
				u.mc, u.ps,
				req.ProgressId,
			)
			if err != nil {
				panic(err)
			}
		})
	}
}

func Import(
	ctx context.Context,
	mc *minio.Client,
	ps stores.ProgressStore,
	progressId uint64,
) error {
	pgs, err := ps.GetProgress(ctx, progressId)
	if err != nil {
		return err
	}
	bytes, err := pgs.TaskDetail.MarshalJSON()
	if err != nil {
		return err
	}
	var detail types.ImportTaskDetail
	err = json.Unmarshal(bytes, &detail)
	if err != nil {
		return err
	}
	sinker := newSinker(detail.DatabaseType)
	if sinker == nil {
		panic("db sinker not implement")
	}

	obj, err := mc.GetObject(
		ctx,
		constants.BucketName,
		detail.FileName,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return err
	}
	extractor, err := newExtractor(detail.FileType, obj)
	if err != nil {
		return err
	}

	defer extractor.Close()

	var buf = make([][]string, 1024)
	index := 0

	for extractor.Next() {
		row, err := extractor.Columns()
		if err != nil {
			return err
		}
		buf[index] = buf[index][:0]
		buf[index] = append(buf[index], row...)
		if index == 1023 {
			err = sinker.Import(ctx,
				detail.Database,
				detail.Table,
				buf,
			)
			if err != nil {
				return err
			}
			buf = buf[:0]
		}
		index = (index + 1) % 1024
	}

	return nil
}

// Presign godoc
// @Summary presign a data file
// @Description presign from oss
// @Tags upload
// @Accept json
// @Produce json
// @Success 200
// @Param req body PresignRequest false "request body"
// @Router /api/v1/presign [get]
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
