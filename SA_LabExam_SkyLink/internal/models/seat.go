package models

import "time"

// Seat ที่นั่งบนอากาศยาน — ระบุตัวตนด้วย "อากาศยานลำใด + หมายเลขที่นั่งใด" ร่วมกัน
// เช่น (aircraft 12, "12A"). หมายเลขที่นั่งจะซ้ำกันได้ข้ามคนละลำ แต่ห้ามซ้ำภายในลำเดียวกัน
type Seat struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	AircraftID uint      `gorm:"column:aircraft_id;not null" json:"aircraft_id"`
	SeatNo     string    `gorm:"column:seat_no;size:4;not null" json:"seat_no"`
	Class      string    `gorm:"size:20;not null;default:economy" json:"class"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Seat) TableName() string {
	return "seats"
}
