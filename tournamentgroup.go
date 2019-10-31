package gotournament

import "fmt"

// TournamentGroupInterface defines the interface of tournament groups used for group tournaments
type TournamentGroupInterface interface {
	GetID() int
	GetTeams() []TeamInterface
	GetGames() []GameInterface
	AppendGames(games []GameInterface)
	AppendGame(game GameInterface)
	AppendTeams(teams []TeamInterface)
	AppendTeam(team TeamInterface)
	Print()
}

// TournamentGroup is for group tournaments only
type TournamentGroup struct {
	ID    int
	Teams []TeamInterface
	Games []GameInterface
}

// GetID returns the id of the group
func (t *TournamentGroup) GetID() int {
	return t.ID
}

// GetTeams returns a slice of teams belonging to the group
func (t *TournamentGroup) GetTeams() []TeamInterface {
	return t.Teams
}

// GetGames returns the slice of games belonging to the group
func (t *TournamentGroup) GetGames() []GameInterface {
	return t.Games
}

// AppendGames adds a slice of games to the Games slice
func (t *TournamentGroup) AppendGames(games []GameInterface) {
	t.Games = append(t.Games, games...)
}

// AppendGame takes a single game and appends it to the Games slice
func (t *TournamentGroup) AppendGame(game GameInterface) {
	t.Games = append(t.Games, game)
}

// AppendTeams adds a slice of teams to the Teams slice
func (t *TournamentGroup) AppendTeams(teams []TeamInterface) {
	t.Teams = append(t.Teams, teams...)
}

// AppendTeam takes a single team and appends it to the Teams slice
func (t *TournamentGroup) AppendTeam(team TeamInterface) {
	t.Teams = append(t.Teams, team)
}

// Print writes group details to stdout
func (t *TournamentGroup) Print() {
	fmt.Printf("Group ID: %d\n", t.GetID())
}
