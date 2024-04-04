package models

type Role struct {
	RoleID   uint   `gorm:"column:role_id;primaryKey;not null"`
	RoleName string `gorm:"column:role_name;not null"`
}
