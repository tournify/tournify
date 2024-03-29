package tournify

import (
	"testing"
)

func TestTournamentGetTypeString(t *testing.T) {
	teamCount := 8
	meetCount := 2
	groupCount := 2
	tournament := CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	if tournament.GetTypeString() != "Group" {
		t.Errorf("Tournament type is not of type Group")
	}
}

func TestTournamentInvalidGroupCount(t *testing.T) {
	teamCount := 8
	meetCount := 2
	groupCount := -1
	tournament := CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	if tournament != nil {
		t.Errorf("Tournament should have failed to create and have a value of nil")
	}
}

func TestTournamentInvalidMeetCount(t *testing.T) {
	teamCount := 8
	meetCount := -1
	groupCount := 2
	tournament := CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	if tournament != nil {
		t.Errorf("Tournament should have failed to create and have a value of nil")
	}
}

func TestCreateGroupTournamentVariation1(t *testing.T) {
	// Test a group tournament with 8 teams in twp groups that meet each other twice
	teamCount := 8
	meetCount := 2
	groupCount := 2
	tournament := CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	assertGroupTournament(t, tournament, teamCount, meetCount, groupCount)
}

func TestCreateGroupTournamentVariation2(t *testing.T) {
	// Test a group tournament with 2 teams that meet once
	teamCount := 2
	meetCount := 1
	groupCount := 1
	tournament := CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))

	assertGroupTournament(t, tournament, teamCount, meetCount, groupCount)
}

func TestCreateGroupTournamentVariation3(t *testing.T) {
	// Test a group tournament with 16 teams in 2 groups that meet each other 4 times
	teamCount := 16
	meetCount := 4
	groupCount := 2
	tournament := CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	assertGroupTournament(t, tournament, teamCount, meetCount, groupCount)
}

func TestCreateGroupTournamentVariation4(t *testing.T) {
	// Test a group tournament with 21 teams in 4 groups that meet each other 1 times
	teamCount := 21
	meetCount := 1
	groupCount := 4
	tournament := CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	assertGroupTournament(t, tournament, teamCount, meetCount, groupCount)
}

func TestCreateGroupTournamentVariation5(t *testing.T) {
	// Test a group tournament with 18 teams in 2 groups that meet each other 4 times
	teamCount := 18
	meetCount := 4
	groupCount := 2
	tournament := CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	assertGroupTournament(t, tournament, teamCount, meetCount, groupCount)
}

func TestCreateGroupTournamentVariation6(t *testing.T) {
	// Test a group tournament with 7 teams in 2 groups that meet each other 1 time
	teamCount := 7
	meetCount := 1
	groupCount := 2
	tournament := CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	assertGroupTournament(t, tournament, teamCount, meetCount, groupCount)
}

func TestCreateGroupTournamentVariation7(t *testing.T) {
	// Test a group tournament with 33 teams in 2 groups that meet each other 1 time
	teamCount := 33
	meetCount := 1
	groupCount := 2
	tournament := CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	assertGroupTournament(t, tournament, teamCount, meetCount, groupCount)
}

func TestCreateGroupTournamentVariation8(t *testing.T) {
	// Test a group tournament with 2 teams in 1 group that meet each other 1 time
	teamCount := 2
	meetCount := 1
	groupCount := 1
	tournament := CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))
	assertGroupTournament(t, tournament, teamCount, meetCount, groupCount)
}

// assertGroupTournament ensures that a group tournament has been created as expected
func assertGroupTournament(t *testing.T, tournament TournamentInterface, teamCount int, meetCount int, groupCount int) {
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

	for _, group := range tournament.GetGroups() {
		for _, groupGame := range *group.GetGames() {
			found := false
			for _, game := range tournament.GetGames() {
				if game.GetID() == groupGame.GetID() {
					found = true
				}
			}
			if !found {
				t.Errorf("A game is present in a group which is not present in the tournament: %d", groupGame.GetID())
			}
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

func TestCreateEliminationTournament(t *testing.T) {
	// Test a group tournament with 8 teams in twp groups that meet each other twice
	teamCount := 8
	tournament := CreateTournament(teamCount, 0, 0, int(TournamentTypeElimination))
	assertEliminationTournament(t, tournament, teamCount)

	teamCount = 9
	tournament = CreateTournament(teamCount, 0, 0, int(TournamentTypeElimination))
	assertEliminationTournament(t, tournament, teamCount)
}

// assertEliminationTournament ensures that a group tournament has been created as expected
func assertEliminationTournament(t *testing.T, tournament TournamentInterface, teamCount int) {
	if len(tournament.GetTeams()) != teamCount {
		t.Errorf("Team count does not match input")
	}
	if len(tournament.GetGames()) != NumberOfGamesForEliminationTournament(teamCount) {
		t.Errorf("Game count for team count %d does not match NumberOfGames calculation: %d != %d", teamCount, len(tournament.GetGames()), NumberOfGamesForEliminationTournament(teamCount))
	}
	for _, game := range tournament.GetGames() {
		if (game.GetHomeTeam().GetID() != -1 && game.GetAwayTeam().GetID() != -1) && game.GetHomeTeam().GetID() == game.GetAwayTeam().GetID() {
			t.Errorf("Two teams with the same ID are facing each other %d vs %d", game.GetHomeTeam().GetID(), game.GetAwayTeam().GetID())
		}
	}
}
