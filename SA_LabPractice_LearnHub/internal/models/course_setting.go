package models

import "time"

// CourseSetting ชุดตั้งค่าของคอร์ส — แต่ละคอร์สมีชุดตั้งค่าได้ไม่เกิน 1 ชุด (1:1)
// ข้อความประกาศ (announcement) อาจไม่ระบุก็ได้
type CourseSetting struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	CourseID     uint   `gorm:"column:course_id;not null" json:"course_id"`
	AllowPreview bool   `gorm:"not null;default:false" json:"allow_preview"`
	Announcement string `gorm:"column:announcement;size:1000;not null" json:"announcement"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (CourseSetting) TableName() string {
	return "course_settings"
}
