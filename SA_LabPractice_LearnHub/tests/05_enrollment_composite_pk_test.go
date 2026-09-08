package tests

import (
	"strings"
	"testing"

	"learnhub/internal/models"
)

// 5. Enrollment is a junction table whose primary key is (learner_id, course_id)
func TestEnrollmentCompositePK(t *testing.T) {
	sch := parseModelSchema(t, &models.Enrollment{})

	if sch.Table != "enrollments" {
		t.Fatalf("expected table name 'enrollments', got '%s'", sch.Table)
	}

	pk := make([]string, 0, len(sch.PrimaryFields))
	for _, f := range sch.PrimaryFields {
		pk = append(pk, strings.ToLower(f.DBName))
	}

	if len(pk) != 2 {
		t.Fatalf("Enrollment must have a composite primary key of exactly 2 columns, got %d: %v", len(pk), pk)
	}

	has := map[string]bool{}
	for _, n := range pk {
		has[n] = true
	}
	if !has["learner_id"] || !has["course_id"] {
		t.Fatalf("Enrollment primary key must be (learner_id, course_id), got %v", pk)
	}
}
