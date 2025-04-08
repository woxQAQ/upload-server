package config

import (
	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"github.com/woxQAQ/upload-server/internal/models"
	"github.com/woxQAQ/upload-server/pkg/constants"
	"gorm.io/gorm"
)

func init() {
	viper.MustBindEnv(constants.DatabaseDSN)
}

func InitDb() *gorm.DB {
	dsn := viper.GetString(constants.DatabaseDSN)
	if dsn == "" {
		panic("dsn missing")
	}

	driver, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	driver.AutoMigrate(models.ImportTaskDetail{}, models.Progress{})
	return driver
}
