package models

import "time"

// Enrollment การลงทะเบียนเรียน — เป็นตารางเชื่อม M:N ระหว่าง Learner (User) กับ Course
// ที่ต้องเก็บข้อมูลเพิ่มบนความสัมพันธ์ (ความคืบหน้า, วันที่ลงทะเบียน)
// ผู้เรียนลงทะเบียนคอร์สเดิมซ้ำไม่ได้
type Enrollment struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	LearnerID  uint      `gorm:"column:learner_id;not null" json:"learner_id"`
	CourseID   uint      `gorm:"column:course_id;not null" json:"course_id"`
	Progress   int       `gorm:"not null;default:0" json:"progress"`
	EnrolledAt time.Time `gorm:"autoCreateTime" json:"enrolled_at"`
}

func (Enrollment) TableName() string {
	return "enrollments"
}
