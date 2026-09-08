package tests

import (
	"reflect"
	"testing"

	"learnhub/internal/models"
)

// 2. Category is a self-referencing hierarchy: a category may have one parent or none
func TestCategorySelfReference(t *testing.T) {
	sch := parseModelSchema(t, &models.Category{})

	parent, ok := sch.FieldsByDBName["parent_id"]
	if !ok || parent.DBName == "" {
		t.Fatalf("expected a 'parent_id' column on categories pointing back to categories")
	}

	if parent.NotNull || hasTag(parent.TagSettings, "NOT NULL", "NOTNULL") {
		t.Fatalf("'parent_id' must be nullable: a top-level category has no parent")
	}

	if parent.FieldType.Kind() != reflect.Ptr {
		t.Fatalf("'parent_id' must model NULL in Go (pointer type) so top-level categories can omit it")
	}

	rel, ok := sch.Relationships.Relations["Parent"]
	if !ok {
		t.Fatalf("expected a self-referencing 'Parent' relationship on Category")
	}
	if rel.FieldSchema == nil || rel.FieldSchema.Table != "categories" {
		t.Fatalf("the 'Parent' relationship must reference the categories table itself")
	}
}
