package models

import (
	"math"
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        int            `gorm:"primarykey;auto_increment;" json:"id"` // auto_increment
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (m *Model) IsNew() (value bool) {
	if m.ID == 0 {
		value = true
	}
	return
}

func Round2(num float32) float32 {
	return float32(math.Round(float64(num)*100)) / 100
}
