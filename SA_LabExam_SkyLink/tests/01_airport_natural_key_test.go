package tests

import (
	"reflect"
	"testing"

	"skylink/internal/models"
)

// 1. Airport is keyed by its natural IATA code (a string), not by a surrogate integer id
func TestAirportNaturalKey(t *testing.T) {
	sch := parseModelSchema(t, &models.Airport{})

	pk := pkDBNames(sch)
	if !sameSet(pk, "code") {
		t.Fatalf("Airport primary key must be exactly the 'code' column, got %v", pk)
	}

	if sch.PrimaryFields[0].FieldType.Kind() != reflect.String {
		t.Fatalf("Airport primary key 'code' must be a string (IATA code), got %s", sch.PrimaryFields[0].FieldType.Kind())
	}

	if f, ok := sch.FieldsByDBName["id"]; ok && f.DBName != "" {
		t.Fatalf("Airport must not carry a surrogate 'id' column; the IATA code is the key")
	}
}
