package common

import "time"

type SQLModel struct {
	Id       int        `json:"id" gorm:"column:id;"`
	Status   int        `json:"status" gorm:"column:id;"`
	CreateAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdateAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}
