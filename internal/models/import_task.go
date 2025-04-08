package models

type ImportTaskDetail struct {
	Model

	OrgName            string `gorm:"index:idx_org_user_db"`
	UserName           string `gorm:"index:idx_org_user_db"`
	DatabaseIdentifier string `gorm:"index:idx_org_user_db"`

	Database string
	Table    string

	FileName string
	FileSize string `gorm:"index:idx_size"`
	OSSUrl   string
}
