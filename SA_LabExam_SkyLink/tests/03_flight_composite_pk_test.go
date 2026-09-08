package tests

import (
	"testing"

	"skylink/internal/models"
)

// 3. A flight is identified by (flight_no, service_date) together
func TestFlightCompositeKey(t *testing.T) {
	sch := parseModelSchema(t, &models.Flight{})

	pk := pkDBNames(sch)
	if !sameSet(pk, "flight_no", "service_date") {
		t.Fatalf("Flight primary key must be the composite (flight_no, service_date), got %v", pk)
	}
}
