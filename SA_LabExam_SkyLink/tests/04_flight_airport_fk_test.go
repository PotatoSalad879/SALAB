package tests

import (
	"testing"

	"skylink/internal/models"
)

// 4. A flight references its origin and destination airports BY IATA CODE (not by surrogate id)
func TestFlightAirportForeignKeys(t *testing.T) {
	sch := parseModelSchema(t, &models.Flight{})

	for _, col := range []string{"origin_code", "destination_code"} {
		if f, ok := sch.FieldsByDBName[col]; !ok || f.DBName == "" {
			t.Fatalf("Flight must persist a '%s' foreign key column", col)
		}
	}

	checkRel := func(name string) {
		rel, ok := sch.Relationships.Relations[name]
		if !ok {
			t.Fatalf("Flight must declare an '%s' relationship to Airport", name)
		}
		if rel.FieldSchema == nil || rel.FieldSchema.Table != "airports" {
			t.Fatalf("Flight.%s must reference the airports table", name)
		}
		if len(rel.References) == 0 {
			t.Fatalf("Flight.%s has no resolved foreign-key reference", name)
		}
		for _, r := range rel.References {
			if r.PrimaryKey == nil || r.PrimaryKey.DBName != "code" {
				pk := "<nil>"
				if r.PrimaryKey != nil {
					pk = r.PrimaryKey.DBName
				}
				t.Fatalf("Flight.%s must reference Airport.code (the natural key), but resolves to airports.%s", name, pk)
			}
		}
	}

	checkRel("Origin")
	checkRel("Destination")
}
