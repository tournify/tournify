package gotournament

type TournamentGroupInterface interface {
	GetID() int
	GetTeams() []TeamInterface
	GetGames() []GameInterface
	AppendGames(games []GameInterface)
	AppendGame(game GameInterface)
	AppendTeams(teams []TeamInterface)
	AppendTeam(team TeamInterface)
}

// TournamentGroup is for group tournaments only
type TournamentGroup struct {
	ID    int
	Teams []TeamInterface
	Games []GameInterface
}

func (t *TournamentGroup) GetID() int {
	return t.ID
}

func (t *TournamentGroup) GetTeams() []TeamInterface {
	return t.Teams
}

func (t *TournamentGroup) GetGames() []GameInterface {
	return t.Games
}

func (t *TournamentGroup) AppendGames(games []GameInterface) {
	t.Games = append(t.Games, games...)
}

func (t *TournamentGroup) AppendGame(game GameInterface) {
	t.Games = append(t.Games, game)
}

func (t *TournamentGroup) AppendTeams(teams []TeamInterface) {
	t.Teams = append(t.Teams, teams...)
}

func (t *TournamentGroup) AppendTeam(team TeamInterface) {
	t.Teams = append(t.Teams, team)
}
