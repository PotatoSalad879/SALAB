package models

import "time"

// Passenger ผู้โดยสาร — ระบุตัวตนด้วยเลขหนังสือเดินทาง
//
// ผู้โดยสารสามารถระบุ "ผู้ร่วมเดินทาง (companions)" ซึ่งก็เป็นผู้โดยสารในระบบเช่นกันได้หลายคน
// และผู้โดยสารหนึ่งคนก็เป็นผู้ร่วมเดินทางของคนอื่นได้หลายคน (ความสัมพันธ์ M:N ที่อ้างอิงตารางตัวเอง)
// ผ่านตารางเชื่อมชื่อ passenger_companions (คอลัมน์ passenger_id กับ companion_id)
type Passenger struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	PassportNo string `gorm:"size:20;not null;uniqueIndex" json:"passport_no"`
	FullName   string `gorm:"size:150;not null" json:"full_name"`

	Companions []Passenger `gorm:"-" json:"companions,omitempty"`

	Notes []Note `gorm:"-" json:"notes,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Passenger) TableName() string {
	return "passengers"
}
