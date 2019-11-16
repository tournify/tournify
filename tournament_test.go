package gotournament

import (
	"testing"
)

func TestCreateGroupTournament(t *testing.T) {
	// Test a group tournament with 8 teams in twp groups that meet each other twice
	teamCount := 8
	meetCount := 2
	groupCount := 2
	tournament := CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	if len(tournament.GetTeams()) != teamCount {
		t.Errorf("Team count does not match input")
	}
	if len(tournament.GetGames()) != NumberOfGamesForGroupTournament(teamCount, groupCount, meetCount) {
		t.Errorf("Game count does not match NumberOfGames calculation: %d != %d", len(tournament.GetGames()), NumberOfGamesForGroupTournament(teamCount, groupCount, meetCount))
	}
	for _, game := range tournament.GetGames() {
		if game.GetHomeTeam().GetID() == game.GetAwayTeam().GetID() {
			t.Errorf("Two teams with the same ID are facing each other %d vs %d", game.GetHomeTeam().GetID(), game.GetAwayTeam().GetID())
		}
	}
	// Test a group tournament with 2 teams that meet once
	teamCount = 2
	meetCount = 1
	groupCount = 1
	tournament = CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	if len(tournament.GetTeams()) != teamCount {
		t.Errorf("Team count does not match input")
	}
	if len(tournament.GetGames()) != NumberOfGamesForGroupTournament(teamCount, groupCount, meetCount) {
		t.Errorf("Game count does not match NumberOfGames calculation: %d != %d", len(tournament.GetGames()), NumberOfGamesForGroupTournament(teamCount, groupCount, meetCount))
	}
	for _, game := range tournament.GetGames() {
		if game.GetHomeTeam().GetID() == game.GetAwayTeam().GetID() {
			t.Errorf("Two teams with the same ID are facing each other %d vs %d", game.GetHomeTeam().GetID(), game.GetAwayTeam().GetID())
		}
	}
	// Test a group tournament with 16 teams in 2 groups that meet each other 4 times
	teamCount = 16
	meetCount = 4
	groupCount = 2
	tournament = CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	if len(tournament.GetTeams()) != teamCount {
		t.Errorf("Team count does not match input")
	}
	if len(tournament.GetGames()) != NumberOfGamesForGroupTournament(teamCount, groupCount, meetCount) {
		t.Errorf("Game count does not match NumberOfGames calculation: %d != %d", len(tournament.GetGames()), NumberOfGamesForGroupTournament(teamCount, groupCount, meetCount))
	}
	for _, game := range tournament.GetGames() {
		if game.GetHomeTeam().GetID() == game.GetAwayTeam().GetID() {
			t.Errorf("Two teams with the same ID are facing each other %d vs %d", game.GetHomeTeam().GetID(), game.GetAwayTeam().GetID())
		}
	}
	// Test a group tournament with 320 teams in 16 groups that meet each other 40 times
	teamCount = 320
	meetCount = 40
	groupCount = 16
	tournament = CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	if len(tournament.GetTeams()) != teamCount {
		t.Errorf("Team count does not match input")
	}
	if len(tournament.GetGames()) != NumberOfGamesForGroupTournament(teamCount, groupCount, meetCount) {
		t.Errorf("Game count does not match NumberOfGames calculation: %d != %d", len(tournament.GetGames()), NumberOfGamesForGroupTournament(teamCount, groupCount, meetCount))
	}
	for _, game := range tournament.GetGames() {
		if game.GetHomeTeam().GetID() == game.GetAwayTeam().GetID() {
			t.Errorf("Two teams with the same ID are facing each other %d vs %d", game.GetHomeTeam().GetID(), game.GetAwayTeam().GetID())
		}
	}
	// Test a group tournament with 18 teams in 2 groups that meet each other 4 times
	teamCount = 18
	meetCount = 4
	groupCount = 2
	tournament = CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	if len(tournament.GetTeams()) != teamCount {
		t.Errorf("Team count does not match input")
	}
	if len(tournament.GetGames()) != NumberOfGamesForGroupTournament(teamCount, groupCount, meetCount) {
		t.Errorf("Game count does not match NumberOfGames calculation: %d != %d", len(tournament.GetGames()), NumberOfGamesForGroupTournament(teamCount, groupCount, meetCount))
	}
	for _, game := range tournament.GetGames() {
		if game.GetHomeTeam().GetID() == game.GetAwayTeam().GetID() {
			t.Errorf("Two teams with the same ID are facing each other %d vs %d", game.GetHomeTeam().GetID(), game.GetAwayTeam().GetID())
		}
	}

	// Test a group tournament with 7 teams in 2 groups that meet each other 1 time
	teamCount = 7
	meetCount = 1
	groupCount = 2
	tournament = CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	if len(tournament.GetTeams()) != teamCount {
		t.Errorf("Team count does not match input")
	}
	if len(tournament.GetGames()) != NumberOfGamesForGroupTournament(teamCount, groupCount, meetCount) {
		t.Errorf("Game count does not match NumberOfGames calculation: %d != %d", len(tournament.GetGames()), NumberOfGamesForGroupTournament(teamCount, groupCount, meetCount))
	}
	for _, game := range tournament.GetGames() {
		if game.GetHomeTeam().GetID() == game.GetAwayTeam().GetID() {
			t.Errorf("Two teams with the same ID are facing each other %d vs %d", game.GetHomeTeam().GetID(), game.GetAwayTeam().GetID())
		}
	}

	// Test a group tournament with 33 teams in 2 groups that meet each other 1 time
	teamCount = 33
	meetCount = 1
	groupCount = 2
	tournament = CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	if len(tournament.GetTeams()) != teamCount {
		t.Errorf("Team count does not match input")
	}
	if len(tournament.GetGames()) != NumberOfGamesForGroupTournament(teamCount, groupCount, meetCount) {
		t.Errorf("Game count does not match NumberOfGames calculation: %d != %d", len(tournament.GetGames()), NumberOfGamesForGroupTournament(teamCount, groupCount, meetCount))
	}
	for _, game := range tournament.GetGames() {
		if game.GetHomeTeam().GetID() == game.GetAwayTeam().GetID() {
			t.Errorf("Two teams with the same ID are facing each other %d vs %d", game.GetHomeTeam().GetID(), game.GetAwayTeam().GetID())
		}
	}

	// Test a group tournament with 2 teams in 1 group that meet each other 1 time
	teamCount = 2
	meetCount = 1
	groupCount = 1
	tournament = CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	if len(tournament.GetTeams()) != teamCount {
		t.Errorf("Team count does not match input")
	}
	if len(tournament.GetGames()) != NumberOfGamesForGroupTournament(teamCount, groupCount, meetCount) {
		t.Errorf("Game count does not match NumberOfGames calculation: %d != %d", len(tournament.GetGames()), NumberOfGamesForGroupTournament(teamCount, groupCount, meetCount))
	}
	for _, game := range tournament.GetGames() {
		if game.GetHomeTeam().GetID() == game.GetAwayTeam().GetID() {
			t.Errorf("Two teams with the same ID are facing each other %d vs %d", game.GetHomeTeam().GetID(), game.GetAwayTeam().GetID())
		}
	}
}

func TestNumberOfGamesForGroupTournament(t *testing.T) {
	// Assert that a tournament with 2 teams split into 2 groups should result in 0 games
	if NumberOfGamesForGroupTournament(2, 2, 1) != 0 {
		t.Errorf("NumberOfGames %d %s", NumberOfGamesForGroupTournament(2, 2, 1), "!= 0")
	}
	// Assert that a tournament with 2 teams in 1 group should result in 1 game
	if NumberOfGamesForGroupTournament(2, 1, 1) != 1 {
		t.Errorf("NumberOfGames %d %s", NumberOfGamesForGroupTournament(2, 1, 1), "!= 1")
	}
	// Assert that a tournament with 4 teams in 1 group should result in 6 games
	if NumberOfGamesForGroupTournament(4, 1, 1) != 6 {
		t.Errorf("NumberOfGames %d %s", NumberOfGamesForGroupTournament(4, 1, 1), "!= 6")
	}
	// Assert that a tournament with 8 teams in 1 group should result in 28 games
	if NumberOfGamesForGroupTournament(8, 1, 1) != 28 {
		t.Errorf("NumberOfGames %d %s", NumberOfGamesForGroupTournament(8, 1, 1), "!= 28")
	}
	// Assert that a tournament with 8 teams split into 2 groups should result in 12 games
	if NumberOfGamesForGroupTournament(8, 2, 1) != 12 {
		t.Errorf("NumberOfGames %d %s", NumberOfGamesForGroupTournament(8, 2, 1), "!= 12")
	}
	// Assert that a tournament with 8 teams split into 2 groups where every team meets 2 times should result in 24 games
	if NumberOfGamesForGroupTournament(8, 2, 2) != 24 {
		t.Errorf("NumberOfGames %d %s", NumberOfGamesForGroupTournament(8, 2, 2), "!= 24")
	}
	// Assert that a tournament with 8 teams split into 2 groups where every team meets 4 times should result in 24 games
	if NumberOfGamesForGroupTournament(8, 2, 4) != 48 {
		t.Errorf("NumberOfGames %d %s", NumberOfGamesForGroupTournament(8, 2, 4), "!= 48")
	}
	// Assert that a tournament with 8 teams split into 2 groups where every team meets 4 times should result in 24 games
	if NumberOfGamesForGroupTournament(16, 2, 4) != 224 {
		t.Errorf("NumberOfGames %d %s", NumberOfGamesForGroupTournament(16, 2, 4), "!= 224")
	}
}
