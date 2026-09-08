package tests

import (
	"testing"

	"skylink/internal/models"
)

// 9. A cancelled booking is kept via soft delete; the booking reference stays unique
func TestBookingSoftDelete(t *testing.T) {
	sch := parseModelSchema(t, &models.Booking{})

	del, ok := sch.FieldsByDBName["deleted_at"]
	if !ok || del.DBName == "" {
		t.Fatalf("Booking must support soft delete (a 'deleted_at' column, e.g. gorm.DeletedAt)")
	}

	ref, ok := sch.FieldsByDBName["ref"]
	if !ok || ref.DBName == "" {
		t.Fatalf("Booking must have a 'ref' column")
	}
	unique := ref.Unique || ref.UniqueIndex != "" || hasTag(ref.TagSettings, "UNIQUE", "UNIQUEINDEX")
	if !unique {
		t.Fatalf("Booking.ref must be unique")
	}
}
