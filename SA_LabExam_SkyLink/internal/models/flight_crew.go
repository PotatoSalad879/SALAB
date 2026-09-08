package models

import "time"

// FlightCrew ตารางมอบหมายลูกเรือเข้าเที่ยวบิน — เป็นตารางเชื่อม M:N ระหว่าง Flight กับ CrewMember
// ที่ต้องเก็บข้อมูลเพิ่มบนความสัมพันธ์ (หน้าที่ในเที่ยวบินนั้น เช่น "captain", "first_officer", "purser")
//
// กุญแจหลักของการมอบหมายคือ "เที่ยวบินใด (หมายเลข + วันที่) + ลูกเรือคนใด" ร่วมกัน
// ลูกเรือคนเดิมถูกมอบหมายเข้าเที่ยวบินเดิมซ้ำไม่ได้
type FlightCrew struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	FlightNo     string    `gorm:"column:flight_no;size:6;not null" json:"flight_no"`
	ServiceDate  time.Time `gorm:"column:service_date;not null" json:"service_date"`
	CrewMemberID uint      `gorm:"column:crew_member_id;not null" json:"crew_member_id"`
	Duty         string    `gorm:"size:20;not null" json:"duty"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (FlightCrew) TableName() string {
	return "flight_crew"
}
