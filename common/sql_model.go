package common

import "time"

type SQLModel struct {
	Id       int       `json:"-" gorm:"column:id;"`
	FakeId   *UID      `json:"id" gorm:"-"`
	Status   int       `json:"status" gorm:"column:status; default: 1"`
	CreateAt time.Time `json:"created_at" gorm:"column:created_at; default:CURRENT_TIMESTAMP()"`
	UpdateAt time.Time `json:"updated_at" gorm:"column:updated_at; default:CURRENT_TIMESTAMP()"`
}

func (m *SQLModel) Mask(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, 1)
	m.FakeId = &uid
}
