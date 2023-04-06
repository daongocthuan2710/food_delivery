package common

import "time"

type SQLModel struct {
	Id       int       `json:"id" gorm:"column:id;"`
	Status   int       `json:"status" gorm:"column:status; default: 1"`
	CreateAt time.Time `json:"created_at" gorm:"column:created_at; default:CURRENT_TIMESTAMP()"`
	UpdateAt time.Time `json:"updated_at" gorm:"column:updated_at; default:CURRENT_TIMESTAMP()"`
}
