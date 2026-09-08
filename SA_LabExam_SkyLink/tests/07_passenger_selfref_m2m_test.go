package tests

import (
	"testing"

	"skylink/internal/models"
	"gorm.io/gorm/schema"
)

// 7. "Companions" is a self-referencing many-to-many on passengers through passenger_companions
func TestPassengerCompanionsSelfM2M(t *testing.T) {
	sch := parseModelSchema(t, &models.Passenger{})

	rel, ok := sch.Relationships.Relations["Companions"]
	if !ok {
		t.Fatalf("Passenger must declare a 'Companions' relationship")
	}
	if rel.Type != schema.Many2Many {
		t.Fatalf("Passenger.Companions must be many2many, got %q", rel.Type)
	}
	if rel.FieldSchema == nil || rel.FieldSchema.Table != "passengers" {
		t.Fatalf("Passenger.Companions must reference the passengers table itself (self-referencing)")
	}
	if rel.JoinTable == nil || rel.JoinTable.Table != "passenger_companions" {
		got := "<nil>"
		if rel.JoinTable != nil {
			got = rel.JoinTable.Table
		}
		t.Fatalf("Passenger.Companions must use join table 'passenger_companions', got %q", got)
	}

	// a self-referencing join must use two DIFFERENT columns for the two sides
	if _, ok := rel.JoinTable.FieldsByDBName["passenger_id"]; !ok {
		t.Fatalf("join table 'passenger_companions' must have a 'passenger_id' column")
	}
	if _, ok := rel.JoinTable.FieldsByDBName["companion_id"]; !ok {
		t.Fatalf("join table 'passenger_companions' must have a distinct 'companion_id' column for the other side (set joinForeignKey / joinReferences)")
	}
}
