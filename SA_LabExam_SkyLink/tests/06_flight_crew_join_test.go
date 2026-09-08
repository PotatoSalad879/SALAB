package tests

import (
	"testing"

	"skylink/internal/models"
)

// 6. FlightCrew is the M:N assignment table carrying an extra column (duty).
//    Its key is (flight_no, service_date, crew_member_id) together.
func TestFlightCrewJoinKey(t *testing.T) {
	sch := parseModelSchema(t, &models.FlightCrew{})

	if sch.Table != "flight_crew" {
		t.Fatalf("expected table 'flight_crew', got '%s'", sch.Table)
	}

	pk := pkDBNames(sch)
	if !sameSet(pk, "flight_no", "service_date", "crew_member_id") {
		t.Fatalf("FlightCrew primary key must be (flight_no, service_date, crew_member_id), got %v", pk)
	}

	if _, ok := sch.FieldsByDBName["duty"]; !ok {
		t.Fatalf("FlightCrew must keep its extra 'duty' column")
	}
}
