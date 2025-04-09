package types

import "net/url"

const (
	Mysql      = "mysql"
	TaskUpload = "upload"
)

type ImportTaskDetail struct {
	DatabaseType string   `json:"databaseType"`
	Database     string   `json:"db"`
	Table        string   `json:"table"`
	FileName     string   `json:"fileName"`
	FileSize     string   `json:"fileSize"`
	OSSUrl       *url.URL `json:"ossUrl,omitempty"`
	MD5          string   `json:"md5"`
}
