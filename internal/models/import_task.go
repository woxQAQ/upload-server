package models

import "gorm.io/gorm"

type ImportTaskDetail struct {
	gorm.Model

	OrgName            string
	UserName           string
	DatabaseIdentifier string

	Database string
	Table    string

	FileName string
	FileSize string
	OSSUrl   string
}
