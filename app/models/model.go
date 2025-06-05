package models

import "gorm.io/plugin/soft_delete"

type Model struct {
	ID        uint  `gorm:"primaryKey"`
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}

type SoftDelete struct {
	DeletedAt soft_delete.DeletedAt
}
