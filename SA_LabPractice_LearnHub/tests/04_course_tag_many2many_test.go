package tests

import (
	"testing"

	"learnhub/internal/models"
	"gorm.io/gorm/schema"
)

// 4. Course and Tag form a many-to-many relationship through a single shared join table
func TestCourseTagManyToMany(t *testing.T) {
	courseSch := parseModelSchema(t, &models.Course{})
	tagSch := parseModelSchema(t, &models.Tag{})

	const wantJoin = "course_tags"

	cRel, ok := courseSch.Relationships.Relations["Tags"]
	if !ok {
		t.Fatalf("Course must declare a 'Tags' relationship to Tag")
	}
	if cRel.Type != schema.Many2Many {
		t.Fatalf("Course.Tags must be a many2many relationship, got %q", cRel.Type)
	}
	if cRel.JoinTable == nil || cRel.JoinTable.Table != wantJoin {
		got := "<nil>"
		if cRel.JoinTable != nil {
			got = cRel.JoinTable.Table
		}
		t.Fatalf("Course.Tags must use join table %q, got %q", wantJoin, got)
	}

	tRel, ok := tagSch.Relationships.Relations["Courses"]
	if !ok {
		t.Fatalf("Tag must declare a 'Courses' relationship to Course")
	}
	if tRel.Type != schema.Many2Many {
		t.Fatalf("Tag.Courses must be a many2many relationship, got %q", tRel.Type)
	}
	if tRel.JoinTable == nil || tRel.JoinTable.Table != wantJoin {
		got := "<nil>"
		if tRel.JoinTable != nil {
			got = tRel.JoinTable.Table
		}
		t.Fatalf("Tag.Courses must use the SAME join table %q as Course.Tags, got %q", wantJoin, got)
	}
}
