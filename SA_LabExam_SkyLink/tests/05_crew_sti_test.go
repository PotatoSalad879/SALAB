package tests

import (
	"reflect"
	"testing"

	"skylink/internal/models"
)

// 5. CrewMember is a Single Table Inheritance model; role-specific fields must be nullable
func TestCrewMemberSTI(t *testing.T) {
	sch := parseModelSchema(t, &models.CrewMember{})

	if sch.Table != "crew_members" {
		t.Fatalf("expected table 'crew_members', got '%s'", sch.Table)
	}

	discriminators := []string{"crew_role", "role", "type"}
	found := false
	for _, d := range discriminators {
		if f, ok := sch.FieldsByDBName[d]; ok && f.DBName != "" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected a role discriminator column (e.g. 'crew_role') on crew_members")
	}

	subtypes := map[string]string{
		"license_no":      "pilot license number",
		"medical_expiry":  "pilot medical expiry",
		"cabin_languages": "attendant languages",
		"service_grade":   "attendant service grade",
	}

	for col, desc := range subtypes {
		f, ok := sch.FieldsByDBName[col]
		if !ok || f.DBName == "" {
			t.Fatalf("expected a '%s' column (%s) on crew_members", col, desc)
		}
		if f.NotNull || hasTag(f.TagSettings, "NOT NULL", "NOTNULL") {
			t.Fatalf("'%s' (%s) must be nullable: the other crew role leaves it empty", col, desc)
		}
		if f.FieldType.Kind() != reflect.Ptr {
			t.Fatalf("'%s' (%s) must be a pointer type to model NULL in Go", col, desc)
		}
	}
}
