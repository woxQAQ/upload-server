package types

import "net/url"

const (
	Mysql      = "mysql"
	Postgres   = "postgres"
	TaskUpload = "upload"
)

type ImportTaskDetail struct {
	DSN          string   `json:"dsn"`
	DatabaseType string   `json:"databaseType"`
	Database     string   `json:"db"`
	Table        string   `json:"table"`
	FileName     string   `json:"fileName"`
	FileType     string   `json:"fileType"`
	FileSize     string   `json:"fileSize"`
	OSSUrl       *url.URL `json:"ossUrl,omitempty"`
	MD5          string   `json:"md5"`
}
