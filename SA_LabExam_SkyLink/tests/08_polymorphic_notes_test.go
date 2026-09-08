package tests

import (
	"testing"

	"skylink/internal/models"
)

// 8. Note is attached polymorphically: both Passenger and Booking own many Notes
//    through the shared (owner_id, owner_type) pair.
func TestNotesPolymorphic(t *testing.T) {
	check := func(model interface{}, name, ownerValue string) {
		sch := parseModelSchema(t, model)
		rel, ok := sch.Relationships.Relations["Notes"]
		if !ok {
			t.Fatalf("%s must declare a 'Notes' relationship", name)
		}
		if rel.Polymorphic == nil {
			t.Fatalf("%s.Notes must be a polymorphic association (polymorphic:Owner)", name)
		}
		if rel.Polymorphic.PolymorphicID == nil || rel.Polymorphic.PolymorphicID.DBName != "owner_id" {
			t.Fatalf("%s.Notes polymorphic id column must be 'owner_id'", name)
		}
		if rel.Polymorphic.PolymorphicType == nil || rel.Polymorphic.PolymorphicType.DBName != "owner_type" {
			t.Fatalf("%s.Notes polymorphic type column must be 'owner_type'", name)
		}
		if rel.Polymorphic.Value != ownerValue {
			t.Fatalf("%s.Notes owner_type value must be %q, got %q", name, ownerValue, rel.Polymorphic.Value)
		}
	}

	check(&models.Passenger{}, "Passenger", "passengers")
	check(&models.Booking{}, "Booking", "bookings")
}
