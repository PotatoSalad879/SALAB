package models

import "time"

// Note บันทึกภายในของเจ้าหน้าที่ — สามารถแนบเข้ากับ "เจ้าของ" ได้หลายชนิด
// (แนบกับ Passenger หรือ Booking ก็ได้) แบบ polymorphic association
// โดยเก็บชนิดของเจ้าของไว้ที่คอลัมน์ owner_type และรหัสเจ้าของที่คอลัมน์ owner_id
type Note struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OwnerID   uint      `gorm:"column:owner_id;not null" json:"owner_id"`
	OwnerType string    `gorm:"column:owner_type;size:30;not null" json:"owner_type"`
	Body      string    `gorm:"size:500;not null" json:"body"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Note) TableName() string {
	return "notes"
}
