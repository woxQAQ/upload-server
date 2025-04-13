package config

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/woxQAQ/upload-server/internal/models"
	"github.com/woxQAQ/upload-server/pkg/types"
)

func InitDb(cfg *types.AppConfig) *gorm.DB {
	if cfg.DatabaseDSN == "" {
		panic("dsn missing")
	}

	driver, err := gorm.Open(sqlite.Open(cfg.DatabaseDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	driver.AutoMigrate(
		models.Progress{},
		models.User{},
		models.UserOrg{},
		models.Organization{},
	)
	return driver
}
