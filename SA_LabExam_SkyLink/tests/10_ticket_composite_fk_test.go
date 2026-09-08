package tests

import (
	"testing"

	"skylink/internal/models"
)

// 10. A ticket points at one flight through the composite foreign key (flight_no, service_date)
func TestTicketCompositeForeignKey(t *testing.T) {
	sch := parseModelSchema(t, &models.Ticket{})

	for _, col := range []string{"flight_no", "service_date", "booking_id"} {
		if f, ok := sch.FieldsByDBName[col]; !ok || f.DBName == "" {
			t.Fatalf("Ticket must persist a '%s' column", col)
		}
	}

	rel, ok := sch.Relationships.Relations["Flight"]
	if !ok {
		t.Fatalf("Ticket must declare a 'Flight' relationship")
	}
	if rel.FieldSchema == nil || rel.FieldSchema.Table != "flights" {
		t.Fatalf("Ticket.Flight must reference the flights table")
	}
	if len(rel.References) < 2 {
		t.Fatalf("Ticket.Flight must be a composite foreign key of 2 columns (flight_no, service_date), resolved refs: %d", len(rel.References))
	}

	fkCols := map[string]bool{}
	for _, r := range rel.References {
		if r.ForeignKey != nil {
			fkCols[r.ForeignKey.DBName] = true
		}
	}
	if !fkCols["flight_no"] || !fkCols["service_date"] {
		t.Fatalf("Ticket.Flight composite foreign key must be (flight_no, service_date), got %v", fkCols)
	}
}
