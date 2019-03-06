package gotournament

import (
	"testing"
)

func TestCreateGroupTournament(t *testing.T) {
	teamCount := 8
	meetCount := 2
	tournament := CreateTournament(teamCount, meetCount, TournamentTypeGroup)
	if len(tournament.Games) < 1 {
		t.Errorf("No games were created for the tournament")
	}
	if len(tournament.Teams) != teamCount {
		t.Errorf("Team count does not match input")
	}
	if len(tournament.Games) != NumberOfGames(teamCount, 2, meetCount) {
		t.Errorf("Game count does not match NumberOfGames calculation: %d != %d", len(tournament.Games), NumberOfGames(teamCount, 2, meetCount))
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

func TestGroupTournamentStats(t *testing.T) {
	//teamCount := 8
	//meetCount := 2
	//tournament := CreateTournament(teamCount, meetCount, TournamentTypeGroup)
	//
	//for _, group := range tournament.Groups {
	//	homeScore := 0
	//	awayScore := len(group.Games)
	//	for i, _ := range group.Games {
	//		group.Games[i].SetScore(float64(homeScore), float64(awayScore))
	//		homeScore++
	//		awayScore--
	//		//fmt.Println(group.Games[i].ID, group.Games[i].HomeTeam.ID, group.Games[i].HomeScore.Points, group.Games[i].AwayTeam.ID, group.Games[i].AwayScore.Points)
	//	}
	//}
}
