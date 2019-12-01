package tournify

import (
	"testing"
)

func TestTournamentTypeString(t *testing.T) {
	ttg1 := TournamentTypeGroup
	if ttg1.String() != "Group" {
		t.Errorf("Tournament Type does not match %s != %s", ttg1.String(), "Group")
	}
	ttg2 := TournamentTypeSeries
	if ttg2.String() != "Series" {
		t.Errorf("Tournament Type does not match %s != %s", ttg2.String(), "Series")
	}
	ttg3 := TournamentTypeElimination
	if ttg3.String() != "Elimination" {
		t.Errorf("Tournament Type does not match %s != %s", ttg3.String(), "Elimination")
	}
	ttg4 := TournamentTypeDoubleElimination
	if ttg4.String() != "Double Elimination" {
		t.Errorf("Tournament Type does not match %s != %s", ttg3.String(), "Double Elimination")
	}
}

func TestTournamentTypeInvalidString(t *testing.T) {
	ttg := TournamentType(-1)
	if ttg.String() != "Unknown" {
		t.Errorf("Tournament Type does not match %s != %s", ttg.String(), "Unknown")
	}
}
