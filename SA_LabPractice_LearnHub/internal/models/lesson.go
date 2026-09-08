package models

import "time"

// Lesson บทเรียนภายในคอร์ส — แต่ละคอร์สมีหลายบทเรียน (1:N)
// เรียงลำดับด้วย order_index และภายในคอร์สเดียวกันห้ามมี order_index ซ้ำกัน
type Lesson struct {
	CourseID   uint   `gorm:"column:course_id;uniqueIndex not null" json:"course_id"`
	Title      string `gorm:"size:200;not null" json:"title"`
	OrderIndex int    `gorm:"column:order_index;uniqueIndex not null" json:"order_index"`

	Course Course `gorm:"foreignKey:CourseID" json:"course,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Lesson) TableName() string {
	return "lessons"
}
