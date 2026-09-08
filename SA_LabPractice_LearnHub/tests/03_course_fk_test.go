package tests

import (
	"testing"

	"learnhub/internal/models"
)

// 3. Course owns the foreign keys to its category and its instructor (1:N on both sides)
func TestCourseForeignKeys(t *testing.T) {
	courseSch := parseModelSchema(t, &models.Course{})
	userSch := parseModelSchema(t, &models.User{})
	catSch := parseModelSchema(t, &models.Category{})

	if _, ok := courseSch.FieldsByDBName["category_id"]; !ok {
		t.Fatalf("Course must persist a 'category_id' foreign key column (each course belongs to exactly one category)")
	}
	if _, ok := courseSch.FieldsByDBName["instructor_id"]; !ok {
		t.Fatalf("Course must persist an 'instructor_id' foreign key column referencing the users table")
	}

	// the users / categories tables must NOT carry a course_id (wrong owning side)
	if f, ok := userSch.FieldsByDBName["course_id"]; ok && f.DBName != "" {
		t.Fatalf("the users table must not own a 'course_id' column; Course is the owning side")
	}
	if f, ok := catSch.FieldsByDBName["course_id"]; ok && f.DBName != "" {
		t.Fatalf("the categories table must not own a 'course_id' column; Course is the owning side")
	}
}
