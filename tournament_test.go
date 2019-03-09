package gotournament

import (
	"testing"
)

func TestCreateGroupTournament(t *testing.T) {
	teamCount := 8
	meetCount := 2
	tournament := CreateTournament(teamCount, meetCount, int(TournamentTypeGroup))
	if len(tournament.GetTeams()) < 1 {
		t.Errorf("No games were created for the tournament")
	}
	if len(tournament.GetTeams()) != teamCount {
		t.Errorf("Team count does not match input")
	}
	if len(tournament.GetGames()) != NumberOfGames(teamCount, 2, meetCount) {
		t.Errorf("Game count does not match NumberOfGames calculation: %d != %d", len(tournament.GetGames()), NumberOfGames(teamCount, 2, meetCount))
	}
}

func TestNumberOfGames(t *testing.T) {
	if NumberOfGames(2, 1, 1) != 1 {
		t.Errorf("NumberOfGames %d %s", NumberOfGames(2, 1, 1), "!= 1")
	}
	if NumberOfGames(4, 1, 1) != 6 {
		t.Errorf("NumberOfGames %d %s", NumberOfGames(4, 1, 1), "!= 6")
	}
	if NumberOfGames(8, 1, 1) != 28 {
		t.Errorf("NumberOfGames %d %s", NumberOfGames(8, 1, 1), "!= 28")
	}
	if NumberOfGames(8, 2, 1) != 12 {
		t.Errorf("NumberOfGames %d %s", NumberOfGames(8, 2, 1), "!= 12")
	}
	if NumberOfGames(8, 2, 2) != 24 {
		t.Errorf("NumberOfGames %d %s", NumberOfGames(8, 2, 2), "!= 24")
	}
}
