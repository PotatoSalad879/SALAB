package models

import "time"

// Booking การจอง — หนึ่งการจองเป็นของผู้โดยสาร 1 คน และมีบัตรโดยสาร (Ticket) ได้หลายใบ
// รหัสการจอง (ref) 6 หลักต้องไม่ซ้ำกันในระบบ
//
// เมื่อการจองถูกยกเลิก ระบบต้อง "ไม่ลบทิ้งจริง" แต่ทำเป็นการลบแบบนุ่มนวล (soft delete)
// เพื่อเก็บประวัติไว้ตรวจสอบย้อนหลัง
type Booking struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Ref         string    `gorm:"size:6;not null;uniqueIndex" json:"ref"`
	PassengerID uint      `gorm:"column:passenger_id;not null" json:"passenger_id"`
	Passenger   Passenger `gorm:"foreignKey:PassengerID" json:"passenger,omitempty"`
	Status      string    `gorm:"size:20;not null;default:pending" json:"status"`

	Notes []Note `gorm:"-" json:"notes,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Booking) TableName() string {
	return "bookings"
}
