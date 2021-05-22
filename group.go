package tournify

import (
	"encoding/json"
	"fmt"
)

// GroupInterface defines the interface of tournament groups used for group tournaments
type GroupInterface interface {
	GetID() int
	GetTeams() *[]TeamInterface
	GetGames() *[]GameInterface
	AppendGames(games []GameInterface)
	AppendGame(game GameInterface)
	AppendTeams(teams []TeamInterface)
	AppendTeam(team TeamInterface)
	Print() string
	Marshal() ([]byte, error)
}

// Group is for group tournaments only
type Group struct {
	ID    int
	Teams []TeamInterface
	Games []GameInterface
}

// GetID returns the id of the group
func (t *Group) GetID() int {
	return t.ID
}

// GetTeams returns a slice of teams belonging to the group
func (t *Group) GetTeams() *[]TeamInterface {
	return &t.Teams
}

// GetGames returns the slice of games belonging to the group
func (t *Group) GetGames() *[]GameInterface {
	return &t.Games
}

// AppendGames adds a slice of games to the Games slice
func (t *Group) AppendGames(games []GameInterface) {
	t.Games = append(t.Games, games...)
}

// AppendGame takes a single game and appends it to the Games slice
func (t *Group) AppendGame(game GameInterface) {
	t.Games = append(t.Games, game)
}

// AppendTeams adds a slice of teams to the Teams slice
func (t *Group) AppendTeams(teams []TeamInterface) {
	t.Teams = append(t.Teams, teams...)
}

// AppendTeam takes a single team and appends it to the Teams slice
func (t *Group) AppendTeam(team TeamInterface) {
	t.Teams = append(t.Teams, team)
}

// Print writes group details to stdout
func (t *Group) Print() string {
	p := fmt.Sprintf("Group ID: %d\n", t.GetID())
	for _, team := range *t.GetTeams() {
		p += team.Print()
	}
	p += fmt.Sprintf("\n")
	return p
}

func (t *Group) Marshal() ([]byte, error) {
	group := struct {
		ID    int
		Teams [][]byte
	}{}
	group.ID = t.ID
	for _, team := range *t.GetTeams() {
		tmpTeam, err := team.Marshal()
		if err != nil {
			return nil, err
		}
		group.Teams = append(group.Teams, tmpTeam)
	}
	return json.Marshal(group)
}
