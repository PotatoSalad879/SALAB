package tests

import (
	"reflect"
	"testing"

	"learnhub/internal/models"
)

// 1. Single Table Inheritance schema for User
func TestUserSingleTableInheritance(t *testing.T) {
	sch := parseModelSchema(t, &models.User{})

	if sch.Table != "users" {
		t.Fatalf("expected User table name to be 'users', got '%s'", sch.Table)
	}

	discriminators := []string{"user_role", "role", "type", "user_type"}
	hasDiscriminator := false
	for _, c := range discriminators {
		if f, ok := sch.FieldsByDBName[c]; ok && f.DBName != "" {
			hasDiscriminator = true
			break
		}
	}
	if !hasDiscriminator {
		t.Fatalf("expected a role discriminator column (e.g. 'user_role') on the users table")
	}

	subtypes := []struct {
		candidates []string
		role       string
	}{
		{[]string{"bio"}, "instructor bio"},
		{[]string{"expertise"}, "instructor expertise"},
		{[]string{"headline"}, "learner headline"},
		{[]string{"learning_goal", "goal"}, "learner learning goal"},
	}

	for _, s := range subtypes {
		var f *fieldInfo
		for _, name := range s.candidates {
			if ff, ok := sch.FieldsByDBName[name]; ok && ff.DBName != "" {
				f = &fieldInfo{notNull: ff.NotNull, kind: ff.FieldType.Kind(), tags: ff.TagSettings, name: ff.Name}
				break
			}
		}
		if f == nil {
			t.Fatalf("expected a nullable subtype column for %s (e.g. %s) on the users table", s.role, s.candidates[0])
		}
		if f.notNull || hasTag(f.tags, "NOT NULL", "NOTNULL") {
			t.Fatalf("subtype column '%s' (%s) must be nullable: rows of the other role leave it empty", f.name, s.role)
		}
		if f.kind != reflect.Ptr {
			t.Fatalf("subtype field '%s' (%s) must model NULL in Go (pointer type)", f.name, s.role)
		}
	}
}

type fieldInfo struct {
	notNull bool
	kind    reflect.Kind
	tags    map[string]string
	name    string
}

func hasTag(tags map[string]string, keys ...string) bool {
	for _, k := range keys {
		if _, ok := tags[k]; ok {
			return true
		}
	}
	return false
}
