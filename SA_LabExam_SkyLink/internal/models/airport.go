package models

import "time"

// Airport สนามบิน — ระบุด้วยรหัส IATA 3 ตัวอักษร (เช่น "BKK", "HKT", "CNX")
// ในระบบนี้รหัส IATA คือกุญแจหลักตามธรรมชาติของสนามบิน ไม่มีการใช้เลขรันเป็นกุญแจหลัก
type Airport struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Code      string    `gorm:"size:3;not null;uniqueIndex" json:"code"`
	Name      string    `gorm:"size:150;not null" json:"name"`
	City      string    `gorm:"size:100;not null" json:"city"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Airport) TableName() string {
	return "airports"
}
