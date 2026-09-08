package models

import "time"

// User เก็บผู้ใช้ทั้งหมดของแพลตฟอร์มไว้ในตารางเดียว (Single Table Inheritance)
// แยกบทบาทด้วยคอลัมน์ user_role ได้แก่ "instructor" และ "learner"
type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	FullName string `gorm:"size:150;not null" json:"full_name"`
	Email    string `gorm:"size:150;not null;uniqueIndex" json:"email"`
	UserRole string `gorm:"column:user_role;size:20;not null" json:"user_role"`

	Bio          *string `gorm:"size:500" json:"bio"`
	Expertise    *string `gorm:"size:150" json:"expertise"`
	Headline     *string `gorm:"size:200" json:"headline"`
	LearningGoal *string `gorm:"size:200" json:"learning_goal"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (User) TableName() string {
	return "users"
}
