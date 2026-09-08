package models

import "time"

// Flight เที่ยวบิน — เที่ยวบินหนึ่ง ๆ ระบุด้วย "หมายเลขเที่ยวบิน + วันที่ให้บริการ" ร่วมกัน
// (เช่น "TG110" ของวันที่ 2026-05-01 เป็นคนละเที่ยวกับ "TG110" ของวันที่ 2026-05-02)
//
// แต่ละเที่ยวบินมีสนามบินต้นทางและปลายทางอย่างละ 1 แห่ง โดยอ้างอิงด้วย "รหัส IATA" ของสนามบิน
// (ไม่ใช่เลขรัน) และมีอากาศยานที่ใช้บิน 1 ลำ
type Flight struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	FlightNo        string    `gorm:"column:flight_no;size:6;not null" json:"flight_no"`
	ServiceDate     time.Time `gorm:"column:service_date;not null" json:"service_date"`
	OriginCode      string    `gorm:"-" json:"origin_code"`
	DestinationCode string    `gorm:"-" json:"destination_code"`
	AircraftID      uint      `gorm:"column:aircraft_id;not null" json:"aircraft_id"`
	Aircraft        Aircraft  `gorm:"foreignKey:AircraftID" json:"aircraft,omitempty"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Flight) TableName() string {
	return "flights"
}
