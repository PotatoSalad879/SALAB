package models

import "time"

// Aircraft อากาศยาน — โครงสร้างนี้ถูกต้องแล้ว ไม่ต้องแก้ไข
type Aircraft struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Registration string    `gorm:"size:10;not null;uniqueIndex" json:"registration"`
	Model        string    `gorm:"size:50;not null" json:"model"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Aircraft) TableName() string {
	return "aircrafts"
}
