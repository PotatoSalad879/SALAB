package tests

import (
	"sync"
	"testing"

	"gorm.io/gorm/schema"
)

var schemaCache sync.Map

// parseModelSchema parses a GORM model struct into a schema.Schema without connecting to any database.
func parseModelSchema(t *testing.T, model interface{}) *schema.Schema {
	t.Helper()

	sch, err := schema.Parse(model, &schemaCache, schema.NamingStrategy{})
	if err != nil {
		t.Fatalf("failed to parse GORM schema for model %T: %v", model, err)
	}
	return sch
}
