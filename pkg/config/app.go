package config

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/woxQAQ/upload-server/internal/controller"
	"gorm.io/gorm"
)

func InitApp(db *gorm.DB, cfg *AppConfig) *gin.Engine {
	e := gin.New()

	minioCli, err := minio.New(cfg.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKeyId, cfg.MinioSecretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		panic(err)
	}

	ok, err := minioCli.BucketExists(context.Background(), "file-upload")
	if err != nil {
		panic(err)
	}
	if !ok {
		minioCli.MakeBucket(context.Background(), "file-upload", minio.MakeBucketOptions{})
	}

	apiV1 := e.Group("/api/v1")

	c := controller.NewUploadController(minioCli)
	c.Register(apiV1)

	return e
}
