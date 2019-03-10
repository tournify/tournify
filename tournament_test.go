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
	if len(tournament.GetGames()) != NumberOfGamesForGroupTournament(teamCount, 2, meetCount) {
		t.Errorf("Game count does not match NumberOfGames calculation: %d != %d", len(tournament.GetGames()), NumberOfGamesForGroupTournament(teamCount, 2, meetCount))
	}
	teamCount = 2
	meetCount = 1
	tournament = CreateTournament(teamCount, meetCount, int(TournamentTypeGroup))
	if len(tournament.GetTeams()) < 1 {
		t.Errorf("No games were created for the tournament")
	}
	if len(tournament.GetTeams()) != teamCount {
		t.Errorf("Team count does not match input")
	}
	if len(tournament.GetGames()) != NumberOfGamesForGroupTournament(teamCount, 1, meetCount) {
		t.Errorf("Game count does not match NumberOfGames calculation: %d != %d", len(tournament.GetGames()), NumberOfGamesForGroupTournament(teamCount, 1, meetCount))
	}
}

func TestNumberOfGamesForGroupTournament(t *testing.T) {
	if NumberOfGamesForGroupTournament(2, 2, 1) != 0 {
		t.Errorf("NumberOfGames %d %s", NumberOfGamesForGroupTournament(2, 2, 1), "!= 0")
	}
	if NumberOfGamesForGroupTournament(2, 1, 1) != 1 {
		t.Errorf("NumberOfGames %d %s", NumberOfGamesForGroupTournament(2, 1, 1), "!= 1")
	}
	if NumberOfGamesForGroupTournament(4, 1, 1) != 6 {
		t.Errorf("NumberOfGames %d %s", NumberOfGamesForGroupTournament(4, 1, 1), "!= 6")
	}
	if NumberOfGamesForGroupTournament(8, 1, 1) != 28 {
		t.Errorf("NumberOfGames %d %s", NumberOfGamesForGroupTournament(8, 1, 1), "!= 28")
	}
	if NumberOfGamesForGroupTournament(8, 2, 1) != 12 {
		t.Errorf("NumberOfGames %d %s", NumberOfGamesForGroupTournament(8, 2, 1), "!= 12")
	}
	if NumberOfGamesForGroupTournament(8, 2, 2) != 24 {
		t.Errorf("NumberOfGames %d %s", NumberOfGamesForGroupTournament(8, 2, 2), "!= 24")
	}
}
