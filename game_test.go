package tournify

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

func TestGameGetParentIDs(t *testing.T) {
	s := []int{2, 3}
	game := Game{
		ID:        1,
		ParentIDs: s,
	}
	if !isEqual(game.GetParentIDs(), s) {
		t.Errorf("ParentIDs are not equal %v: %v\n", game.GetParentIDs(), s)
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

	score := Score{Points: 10}
	teams2 := []TeamInterface{&homeTeam, &awayTeam}
	game2 := Game{
		ID:     1,
		Teams:  teams2,
		Scores: []ScoreInterface{&score},
	}
	game2.SetScore(homeScore, awayScore)
	if game2.GetHomeScore().GetPoints() != homeScore {
		t.Errorf("HomeScore does not match %.2f != %.2f\n", game.GetHomeScore().GetPoints(), homeScore)
	}
	if game2.GetAwayScore().GetPoints() != awayScore {
		t.Errorf("AwayScore does not match %.2f != %.2f\n", game.GetAwayScore().GetPoints(), awayScore)
	}
}

func TestGameGetHomeAwayTeams(t *testing.T) {
	homeTeam := Team{ID: 1}
	awayTeam := Team{ID: 2}
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

	game2 := Game{
		ID: 1,
	}
	if game2.GetHomeTeam().GetID() != 0 {
		t.Errorf("Home team ids does not match %d != %d\n", game2.GetHomeTeam().GetID(), 0)
	}
	if game2.GetAwayTeam().GetID() != 0 {
		t.Errorf("Away team ids does not match %d != %d\n", game2.GetAwayTeam().GetID(), 0)
	}
}

func TestGameSetHomeTeam(t *testing.T) {
	homeTeam := Team{ID: 1}
	game := Game{
		ID: 1,
	}

	game.SetHomeTeam(&homeTeam)
	if game.GetHomeTeam().GetID() != homeTeam.GetID() {
		t.Errorf("Home team ids does not match %d != %d\n", game.GetHomeTeam().GetID(), homeTeam.GetID())
	}

	homeTeam2 := Team{ID: 99}
	game.SetHomeTeam(&homeTeam2)
	if game.GetHomeTeam().GetID() != homeTeam2.GetID() {
		t.Errorf("Home team ids does not match %d != %d\n", game.GetHomeTeam().GetID(), homeTeam2.GetID())
	}
}

func TestGameSetAwayTeam(t *testing.T) {
	awayTeam := Team{ID: 1}
	game := Game{
		ID: 1,
	}

	game.SetAwayTeam(&awayTeam)
	if game.GetAwayTeam().GetID() != awayTeam.GetID() {
		t.Errorf("Away team ids does not match %d != %d\n", game.GetAwayTeam().GetID(), awayTeam.GetID())
	}

	if game.GetHomeTeam().GetID() != 0 {
		t.Errorf("Home team ids does not match %d != %d\n", game.GetHomeTeam().GetID(), 0)
	}

	homeTeam := Team{ID: 22}
	game.SetHomeTeam(&homeTeam)

	awayTeam2 := Team{ID: 99}
	game.SetAwayTeam(&awayTeam2)
	if game.GetAwayTeam().GetID() != awayTeam2.GetID() {
		t.Errorf("Away team ids does not match %d != %d\n", game.GetAwayTeam().GetID(), awayTeam2.GetID())
	}
	if game.GetHomeTeam().GetID() != homeTeam.GetID() {
		t.Errorf("Home team ids does not match %d != %d\n", game.GetHomeTeam().GetID(), homeTeam.GetID())
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
