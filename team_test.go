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

func TestTeamGetGames(t *testing.T) {
	game1 := Game{ID: 1}
	game2 := Game{ID: 2}
	game3 := Game{ID: 3}
	game4 := Game{ID: 4}
	team := Team{
		ID: 0,
		Games: []GameInterface{
			&game1,
			&game2,
			&game3,
			&game4,
		},
	}
	if len(team.GetGames()) != 4 {
		t.Errorf("Number of games for team does not match %d != %d", len(team.GetGames()), 4)
	}
	for i, g := range team.GetGames() {
		if i == 0 && g.GetID() != 1 {
			t.Errorf("Game ID is not correct %d != %d", g.GetID(), 1)
		}
		if i == 1 && g.GetID() != 2 {
			t.Errorf("Game ID is not correct %d != %d", g.GetID(), 2)
		}
		if i == 2 && g.GetID() != 3 {
			t.Errorf("Game ID is not correct %d != %d", g.GetID(), 3)
		}
		if i == 3 && g.GetID() != 4 {
			t.Errorf("Game ID is not correct %d != %d", g.GetID(), 4)
		}
	}
}

func TestTeamAppendGame(t *testing.T) {
	game := Game{ID: 0}
	team := Team{ID: 0}
	team.AppendGame(&game)
	if len(team.GetGames()) != 1 {
		t.Errorf("Number of games for team does not match %d != %d", len(team.GetGames()), 4)
	}
	for i, g := range team.GetGames() {
		if i == 0 && g.GetID() != 0 {
			t.Errorf("Game ID is not correct %d != %d", g.GetID(), 0)
		}
	}
}

func TestTeamPrint(t *testing.T) {
	game1 := Game{ID: 1}
	game2 := Game{ID: 2}
	game3 := Game{ID: 3}
	game4 := Game{ID: 4}
	team := Team{
		ID: 0,
		Games: []GameInterface{
			&game1,
			&game2,
			&game3,
			&game4,
		},
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
