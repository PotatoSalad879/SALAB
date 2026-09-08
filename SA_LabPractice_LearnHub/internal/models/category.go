package models

import "time"

// Category หมวดหมู่คอร์ส จัดเป็นลำดับชั้น — หมวดหมู่หนึ่งอาจมีหมวดหมู่แม่ได้ 1 หมวด
// หรือเป็นหมวดระดับบนสุดก็ได้ และหนึ่งหมวดหมู่มีคอร์สได้หลายคอร์ส (1:N)
type Category struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `gorm:"size:120;not null" json:"name"`
	Slug     string    `gorm:"size:120;not null;uniqueIndex" json:"slug"`
	ParentID *uint      `gorm:"" json:"parent_id"`
	Parent   *Category `gorm:"foreignKey:ParentID"`

	Courses []Course `gorm:"foreignKey:CategoryID" json:"courses,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Category) TableName() string {
	return "categories"
}
