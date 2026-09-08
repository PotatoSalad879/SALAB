package tests

import (
	"testing"

	"skylink/internal/models"
)

// 2. A seat is identified by (aircraft_id, seat_no) together
func TestSeatCompositeKey(t *testing.T) {
	sch := parseModelSchema(t, &models.Seat{})

	pk := pkDBNames(sch)
	if !sameSet(pk, "aircraft_id", "seat_no") {
		t.Fatalf("Seat primary key must be the composite (aircraft_id, seat_no), got %v", pk)
	}
}
