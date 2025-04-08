package config

import (
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
		Secure: true,
	})

	if err != nil {
		panic(err)
	}

	apiV1 := e.Group("/api/v1")

	c := controller.NewUploadController(minioCli)
	c.Register(apiV1)

	return e
}
