package models

import "time"

// Ticket บัตรโดยสาร — หนึ่งใบอยู่ใต้การจอง 1 รายการ และผูกกับเที่ยวบิน 1 เที่ยว
// การอ้างอิงเที่ยวบินต้องใช้ "หมายเลขเที่ยวบิน + วันที่ให้บริการ" ร่วมกัน (Foreign Key แบบหลายคอลัมน์)
// ให้ตรงกับกุญแจหลักของ Flight
type Ticket struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	BookingID   uint      `gorm:"column:booking_id;not null" json:"booking_id"`
	Booking     Booking   `gorm:"foreignKey:BookingID" json:"booking,omitempty"`
	FlightNo    string    `gorm:"-" json:"flight_no"`
	ServiceDate time.Time `gorm:"-" json:"service_date"`
	FareAmount  float64   `gorm:"type:numeric(10,2);not null" json:"fare_amount"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Ticket) TableName() string {
	return "tickets"
}
