package config

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gorm.io/gorm"

	"github.com/woxQAQ/upload-server/internal/controller"
	stores "github.com/woxQAQ/upload-server/internal/stores/progress"
	"github.com/woxQAQ/upload-server/pkg/constants"
	"github.com/woxQAQ/upload-server/pkg/storage"
	"github.com/woxQAQ/upload-server/pkg/types"
)

func InitApp(db *gorm.DB, cfg *types.AppConfig) *gin.Engine {
	e := gin.New()

	mc, err := minio.New(cfg.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKeyId, cfg.MinioSecretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	storage.MustCheckBucket(ctx, mc, constants.BucketName)

	apiV1 := e.Group("/api/v1")
	ps := stores.NewProgressStore(db)

	c := controller.NewUploadController(mc, ps)
	c.Register(apiV1)

	return e
}
