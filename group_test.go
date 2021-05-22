package tournify

import (
	"testing"
)

func TestTournamentGroupGetID(t *testing.T) {
	tg := Group{
		ID: 1,
	}
	if tg.GetID() != 1 {
		t.Errorf("Tournament Group ID is not equal to 1\n")
	}
}

func TestTournamentGroupGetTeams(t *testing.T) {
	team1 := Team{ID: 1}
	team2 := Team{ID: 2}
	team3 := Team{ID: 3}
	team4 := Team{ID: 4}
	tg := Group{
		ID: 0,
		Teams: []TeamInterface{
			&team1,
			&team2,
			&team3,
			&team4,
		},
	}
	assertTeams(tg, t)
}

func TestTournamentGroupAppendTeam(t *testing.T) {
	team1 := Team{ID: 1}
	team2 := Team{ID: 2}
	team3 := Team{ID: 3}
	team4 := Team{ID: 4}
	tg := Group{
		ID: 0,
	}
	tg.AppendTeam(&team1)
	tg.AppendTeam(&team2)
	tg.AppendTeam(&team3)
	tg.AppendTeam(&team4)
	assertTeams(tg, t)
}

func TestTournamentGroupAppendTeams(t *testing.T) {
	team1 := Team{ID: 1}
	team2 := Team{ID: 2}
	team3 := Team{ID: 3}
	team4 := Team{ID: 4}
	tg := Group{
		ID: 0,
	}
	tg.AppendTeams([]TeamInterface{&team1, &team2, &team3, &team4})
	assertTeams(tg, t)
}

func TestTournamentGroupGetGames(t *testing.T) {
	game1 := Game{ID: 1}
	game2 := Game{ID: 2}
	game3 := Game{ID: 3}
	game4 := Game{ID: 4}
	tg := Group{
		ID: 0,
		Games: []GameInterface{
			&game1,
			&game2,
			&game3,
			&game4,
		},
	}
	assertGames(tg, t)
}

func TestTournamentGroupAppendGame(t *testing.T) {
	game1 := Game{ID: 1}
	game2 := Game{ID: 2}
	game3 := Game{ID: 3}
	game4 := Game{ID: 4}
	tg := Group{
		ID: 0,
	}
	tg.AppendGame(&game1)
	tg.AppendGame(&game2)
	tg.AppendGame(&game3)
	tg.AppendGame(&game4)
	assertGames(tg, t)
}

func TestTournamentGroupAppendGames(t *testing.T) {
	game1 := Game{ID: 1}
	game2 := Game{ID: 2}
	game3 := Game{ID: 3}
	game4 := Game{ID: 4}
	tg := Group{
		ID: 0,
	}
	tg.AppendGames([]GameInterface{&game1, &game2, &game3, &game4})
	assertGames(tg, t)
}

func TestTournamentGroupPrint(t *testing.T) {
	team1 := Team{ID: 1}
	team2 := Team{ID: 2}
	team3 := Team{ID: 3}
	team4 := Team{ID: 4}
	game1 := Game{ID: 1}
	game2 := Game{ID: 2}
	game3 := Game{ID: 3}
	game4 := Game{ID: 4}
	tg := Group{
		ID: 0,
		Games: []GameInterface{
			&game1,
			&game2,
			&game3,
			&game4,
		},
		Teams: []TeamInterface{
			&team1,
			&team2,
			&team3,
			&team4,
		},
	}
	if tg.Print() != "Group ID: 0\nTeam ID: 1\nTeam ID: 2\nTeam ID: 3\nTeam ID: 4\n\n" {
		t.Errorf("Tournament group print does not match %s != %s", tg.Print(), "Group ID: 0\nTeam ID: 1\nTeam ID: 2\nTeam ID: 3\nTeam ID: 4\n\n")
	}
}

func assertGames(tg Group, t *testing.T) {
	if len(*tg.GetGames()) != 4 {
		t.Errorf("Number of games on tournament group does not match %d != %d", len(*tg.GetGames()), 4)
	}
	for i, game := range *tg.GetGames() {
		if game.GetID() != i+1 {
			t.Errorf("Game ID is not correct %d != %d", game.GetID(), 1)
		}
	}
}

func assertTeams(tg Group, t *testing.T) {
	if len(*tg.GetTeams()) != 4 {
		t.Errorf("Number of teams on tournament group does not match %d != %d", len(*tg.GetTeams()), 4)
	}
	for i, team := range *tg.GetTeams() {
		if team.GetID() != i+1 {
			t.Errorf("Team ID is not correct %d != %d", team.GetID(), 1)
		}
	}
}
