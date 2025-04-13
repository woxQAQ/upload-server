package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/woxQAQ/upload-server/pkg/snowflake"
)

type Model struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (d *Model) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID, err = snowflake.GetSnowFlakeID()
	if err != nil {
		return err
	}
	return
}
