package models

import "time"

// Tag ป้ายกำกับคอร์ส — ใช้ซ้ำกับคอร์สได้หลายคอร์ส (M:N กับ Course ผ่านตารางเชื่อม course_tags)
type Tag struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:80;not null;uniqueIndex" json:"name"`

	Courses []Course `gorm:"many2many:tag_courses" json:"courses,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Tag) TableName() string {
	return "tags"
}
