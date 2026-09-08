package tests

import (
	"reflect"
	"testing"

	"learnhub/internal/models"
)

// 7. A learner may review a given course at most once; the review comment is optional
func TestReviewConstraints(t *testing.T) {
	sch := parseModelSchema(t, &models.Review{})

	for _, col := range []string{"course_id", "learner_id"} {
		if _, ok := sch.FieldsByDBName[col]; !ok {
			t.Fatalf("Review must persist a '%s' column", col)
		}
	}

	if !hasCompositeUniqueIndex(sch, "course_id", "learner_id") {
		t.Fatalf("Review must enforce a composite UNIQUE index on (course_id, learner_id): one learner reviews one course at most once")
	}

	comment, ok := sch.FieldsByDBName["comment"]
	if !ok || comment.DBName == "" {
		t.Fatalf("Review must have a 'comment' column")
	}
	if comment.NotNull || hasTag(comment.TagSettings, "NOT NULL", "NOTNULL") {
		t.Fatalf("'comment' must be nullable: a review may omit free text")
	}
	if comment.FieldType.Kind() != reflect.Ptr {
		t.Fatalf("'comment' must model NULL in Go (pointer type)")
	}
}
