package models

type UserOrg struct {
	Model
	UserName string `gorm:"index:idx_user_orgid_org;not null"`
	OrgId    string `gorm:"index:idx_user_orgid_org;not null"`
	OrgName  string `gorm:"index:idx_user_orgid_org;not null"`
	Role     string `gorm:"index:idx_role;not null"`
}
