package tests

import (
	"strings"
	"testing"

	"learnhub/internal/models"
	"gorm.io/gorm/schema"
)

// hasCompositeUniqueIndex reports whether sch has a UNIQUE index covering exactly the given columns.
func hasCompositeUniqueIndex(sch *schema.Schema, cols ...string) bool {
	want := map[string]bool{}
	for _, c := range cols {
		want[c] = true
	}
	for _, idx := range sch.ParseIndexes() {
		if strings.ToUpper(idx.Class) != "UNIQUE" {
			continue
		}
		got := map[string]bool{}
		for _, f := range idx.Fields {
			got[f.DBName] = true
		}
		if len(got) != len(want) {
			continue
		}
		ok := true
		for c := range want {
			if !got[c] {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}

// 6. Within one course, lesson order_index values must be unique (composite unique index)
func TestLessonOrderUniquePerCourse(t *testing.T) {
	sch := parseModelSchema(t, &models.Lesson{})

	if _, ok := sch.FieldsByDBName["course_id"]; !ok {
		t.Fatalf("Lesson must persist a 'course_id' foreign key column")
	}
	if _, ok := sch.FieldsByDBName["order_index"]; !ok {
		t.Fatalf("Lesson must persist an 'order_index' column")
	}

	if !hasCompositeUniqueIndex(sch, "course_id", "order_index") {
		t.Fatalf("Lesson must enforce a composite UNIQUE index on (course_id, order_index) so no two lessons of the same course share an order")
	}
}
