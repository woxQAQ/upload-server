package controller

type PresignRequest struct {
	FileName string `json:"filename" binding:"required"`
	FileSize string `json:"filesize" binding:"required"`
	Database string `json:"database" binding:"required"`
	Table    string `json:"table" binding:"required"`
	MD5      string `json:"md5" binding:"required"`
}

type ApproveRequest struct {
	ProgressId uint64  `json:"id" binding:"require"`
	Approver   string  `json:"approver" binding:"required"`
	Opinion    *string `json:"opionion"`
	Trigger    string  `json:"trigger" default:"current" binding:"oneof 'current after'"`
}

const (
	TriggerCurrent = "current"
	TriggerAfter   = "after"
)
