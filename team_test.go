package tournify

import (
	"testing"
)

func TestTeamGetID(t *testing.T) {
	team := Team{
		ID: 1,
	}
	if team.GetID() != 1 {
		t.Errorf("Team ID is not equal to 1\n")
	}
}

func TestTeamPrint(t *testing.T) {
	team := Team{
		ID: 0,
	}
	if team.Print() != "Team ID: 0\n" {
		t.Errorf("Team print does not match %s != %s", team.Print(), "Team ID: 0\n")
	}
}

func TestTeamGetElimnatedCount(t *testing.T) {
	team := Team{
		ID:         0,
		Eliminated: 1,
	}
	if team.GetEliminatedCount() != 1 {
		t.Errorf("Team GetEliminatedCount does not match %d != %d", team.GetEliminatedCount(), 1)
	}
}
