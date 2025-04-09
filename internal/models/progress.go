package models

import "encoding/json"

type Progress struct {
	Model
	Name            string          `gorm:"not null;index:idx_name"`
	Kind            string          `gorm:"not null"`
	TaskDetail      json.RawMessage `gorm:"not null"`
	OrgName         string          `gorm:"not null;index:idx_org_user"`
	UserId          string          `gorm:"not null;index:idx_org_user"`
	Status          string          `gorm:"not null;index:idx_status_approver"`
	Description     *string
	Approver        *string `gorm:"index:idx_status_approver"`
	ApprovalOpinion *string
}
