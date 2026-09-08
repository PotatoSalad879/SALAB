package tests

import (
	"reflect"
	"testing"

	"learnhub/internal/models"
)

// 8. Each course has at most one settings row (1:1); the announcement text is optional
func TestCourseSettingOneToOne(t *testing.T) {
	sch := parseModelSchema(t, &models.CourseSetting{})

	if sch.Table != "course_settings" {
		t.Fatalf("expected table name 'course_settings', got '%s'", sch.Table)
	}

	course, ok := sch.FieldsByDBName["course_id"]
	if !ok || course.DBName == "" {
		t.Fatalf("CourseSetting must persist a 'course_id' foreign key column")
	}

	unique := course.Unique || course.UniqueIndex != "" ||
		hasTag(course.TagSettings, "UNIQUE", "UNIQUEINDEX") ||
		hasCompositeUniqueIndex(sch, "course_id")
	if !unique {
		t.Fatalf("'course_id' must be UNIQUE to enforce the 1:1 relationship between Course and CourseSetting")
	}

	ann, ok := sch.FieldsByDBName["announcement"]
	if !ok || ann.DBName == "" {
		t.Fatalf("CourseSetting must have an 'announcement' column")
	}
	if ann.NotNull || hasTag(ann.TagSettings, "NOT NULL", "NOTNULL") {
		t.Fatalf("'announcement' must be nullable: a course may have no announcement")
	}
	if ann.FieldType.Kind() != reflect.Ptr {
		t.Fatalf("'announcement' must model NULL in Go (pointer type)")
	}
}
