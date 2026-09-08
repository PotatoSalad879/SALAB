package models

import "time"

// Course คอร์สเรียน — แต่ละคอร์สอยู่ใน 1 หมวดหมู่ และมีผู้สอนหลัก 1 คน (อ้างไปตาราง users)
// คอร์สหนึ่งติดแท็กได้หลายอัน และแท็กหนึ่งใช้กับคอร์สได้หลายคอร์ส (M:N) ผ่านตารางเชื่อม course_tags
type Course struct {
	ID    uint    `gorm:"primaryKey" json:"id"`
	Title string  `gorm:"size:200;not null" json:"title"`
	Slug  string  `gorm:"size:200;not null;uniqueIndex" json:"slug"`
	Price float64 `gorm:"type:numeric(10,2);not null" json:"price"`

	CategoryID   uint `gorm:"" json:"category_id"`
	InstructorID uint `gorm:"" json:"instructor_id"`

	Category   Category `gorm:"foreignKey:CategoryID"`
	Instructor User     `gorm:"foreignKey:InstructorID"`

	Tags []Tag `gorm:"many2many:course_tags" json:"tags,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Course) TableName() string {
	return "courses"
}
