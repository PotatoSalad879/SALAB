package tests

import (
	"strings"
	"sync"
	"testing"

	"gorm.io/gorm/schema"
)

var schemaCache sync.Map

// parseModelSchema parses a GORM model struct into a schema.Schema without connecting to any database.
func parseModelSchema(t testing.TB, model interface{}) *schema.Schema {
	t.Helper()

	sch, err := schema.Parse(model, &schemaCache, schema.NamingStrategy{})
	if err != nil {
		t.Fatalf("failed to parse GORM schema for model %T: %v", model, err)
	}
	return sch
}

func pkDBNames(sch *schema.Schema) []string {
	out := make([]string, 0, len(sch.PrimaryFields))
	for _, f := range sch.PrimaryFields {
		out = append(out, strings.ToLower(f.DBName))
	}
	return out
}

func sameSet(got []string, want ...string) bool {
	if len(got) != len(want) {
		return false
	}
	m := map[string]bool{}
	for _, g := range got {
		m[g] = true
	}
	for _, w := range want {
		if !m[w] {
			return false
		}
	}
	return true
}

func hasTag(tags map[string]string, keys ...string) bool {
	for _, k := range keys {
		if _, ok := tags[k]; ok {
			return true
		}
	}
	return false
}
