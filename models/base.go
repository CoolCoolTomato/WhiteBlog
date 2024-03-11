package models

import "time"

type BaseModel struct {
	ID          int       `gorm:"primary_key"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"updated_date"`
}
