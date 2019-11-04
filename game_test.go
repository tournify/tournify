package gotournament

import (
	"testing"
)

func TestGameGetID(t *testing.T) {
	game := Game{
		ID: 1,
	}
	if game.GetID() != 1 {
		t.Errorf("Game ID is not equal to 1\n")
	}
}

func TestGameSetScore(t *testing.T) {
	homeScore := 10.0
	awayScore := 5.0
	homeTeam := Team{ID: 0}
	awayTeam := Team{ID: 1}
	teams := []TeamInterface{&homeTeam, &awayTeam}
	game := Game{
		ID:    1,
		Teams: teams,
	}
	game.SetScore(homeScore, awayScore)
	if game.GetHomeScore().GetPoints() != homeScore {
		t.Errorf("HomeScore does not match %.2f != %.2f\n", game.GetHomeScore().GetPoints(), homeScore)
	}
	if game.GetAwayScore().GetPoints() != awayScore {
		t.Errorf("AwayScore does not match %.2f != %.2f\n", game.GetAwayScore().GetPoints(), awayScore)
	}
}

func TestGameGetHomeAwayTeams(t *testing.T) {
	homeTeam := Team{ID: 0}
	awayTeam := Team{ID: 1}
	teams := []TeamInterface{&homeTeam, &awayTeam}
	game := Game{
		ID:    1,
		Teams: teams,
	}
	if game.GetHomeTeam().GetID() != homeTeam.GetID() {
		t.Errorf("Home team ids does not match %d != %d\n", game.GetHomeTeam().GetID(), homeTeam.GetID())
	}
	if game.GetAwayTeam().GetID() != awayTeam.GetID() {
		t.Errorf("Away team ids does not match %d != %d\n", game.GetAwayTeam().GetID(), awayTeam.GetID())
	}
}

func TestGameGetTeams(t *testing.T) {
	homeTeam := Team{ID: 0}
	awayTeam := Team{ID: 1}
	teams := []TeamInterface{&homeTeam, &awayTeam}
	game := Game{
		ID:    1,
		Teams: teams,
	}
	teams2 := game.GetTeams()
	for i, te := range teams {
		if te.GetID() != teams2[i].GetID() {
			t.Errorf("Team ids do not match %d != %d\n", te.GetID(), teams2[i].GetID())
		}
	}
}

func TestGameGetScores(t *testing.T) {
	homeScore := 10.0
	awayScore := 5.0
	homeTeam := Team{ID: 0}
	awayTeam := Team{ID: 1}
	teams := []TeamInterface{&homeTeam, &awayTeam}
	game := Game{
		ID:    1,
		Teams: teams,
	}
	game.SetScore(homeScore, awayScore)
	scores := game.GetScores()
	for i, s := range scores {
		if i == 0 && s.GetPoints() != homeScore {
			t.Errorf("Scores do not match %.2f != %.2f\n", s.GetPoints(), homeScore)
		}
		if i == 1 && s.GetPoints() != awayScore {
			t.Errorf("Scores do not match %.2f != %.2f\n", s.GetPoints(), awayScore)
		}
	}
}

func TestGamePrint(t *testing.T) {
	homeScore := 10.0
	awayScore := 5.0
	homeTeam := Team{ID: 0}
	awayTeam := Team{ID: 1}
	teams := []TeamInterface{&homeTeam, &awayTeam}
	game := Game{
		ID:    1,
		Teams: teams,
	}
	game.SetScore(homeScore, awayScore)
	if game.Print() != "Game ID: 1, HomeTeam: 0, AwayTeam: 1, HomeScore: 10.00, AwayScore: 5.00\n" {
		t.Errorf("Print does not match %s != %s\n", game.Print(), "Game ID: 1, HomeTeam: 0, AwayTeam: 1, HomeScore: 10.00, AwayScore: 5.00\n")
	}
}
