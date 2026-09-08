package models

import "time"

// CrewMember ลูกเรือ — เก็บลูกเรือทุกคนไว้ในตารางเดียว (Single Table Inheritance)
// แยกบทบาทด้วยคอลัมน์ crew_role:
//   - "pilot"     : มีข้อมูลเฉพาะ เลขที่ใบอนุญาตนักบิน (license_no) และวันหมดอายุใบตรวจร่างกาย (medical_expiry)
//   - "attendant" : มีข้อมูลเฉพาะ ภาษาที่ให้บริการได้ (cabin_languages) และระดับการบริการ (service_grade)
// ข้อมูลเฉพาะของแต่ละบทบาทต้องเว้นว่างได้เมื่อบันทึกลูกเรืออีกบทบาทหนึ่ง
type CrewMember struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	StaffNo  string `gorm:"size:10;not null;uniqueIndex" json:"staff_no"`
	FullName string `gorm:"size:150;not null" json:"full_name"`
	CrewRole string `gorm:"column:crew_role;size:20;not null" json:"crew_role"`

	LicenseNo      string `gorm:"size:30;not null" json:"license_no"`
	MedicalExpiry  string `gorm:"size:30;not null" json:"medical_expiry"`
	CabinLanguages string `gorm:"size:120;not null" json:"cabin_languages"`
	ServiceGrade   string `gorm:"size:20;not null" json:"service_grade"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (CrewMember) TableName() string {
	return "crew_members"
}
