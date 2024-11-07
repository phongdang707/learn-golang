package common

import "time"

type SQLModel struct {
	Id          int        `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"column:updated_at"`
}
