package models

import "gorm.io/gorm"

type Progress struct {
	gorm.Model
	Kind        string
	TaskId      string
	OrgName     string `gorm:"index:idx_org_user"`
	User        string `gorm:"index:idx_org_user"`
	Status      string `gorm:"index:idx_status_approver"`
	Description string
	Approver    string
}
