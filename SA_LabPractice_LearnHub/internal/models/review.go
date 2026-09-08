package models

import "time"

// Review รีวิวคอร์สจากผู้เรียน — ผู้เรียน 1 คนรีวิวคอร์สเดิมได้ไม่เกิน 1 ครั้ง
// ข้อความรีวิว (comment) อาจไม่ระบุก็ได้
type Review struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	CourseID  uint   `gorm:"column:course_id;not null" json:"course_id"`
	LearnerID uint   `gorm:"column:learner_id;not null" json:"learner_id"`
	Rating    int    `gorm:"not null" json:"rating"`
	Comment   string `gorm:"column:comment;size:1000;not null" json:"comment"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Review) TableName() string {
	return "reviews"
}
