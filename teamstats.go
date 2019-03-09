package gotournament

type TeamStatsInterface interface {
	GetID() int
	GetTournament() TournamentInterface
	GetGroup() TournamentGroupInterface
	GetTeam() TeamInterface
	GetPlayed() int
	GetWins() int
	GetLosses() int
	GetTies() int
	GetPointsFor() float64
	GetPointsAgainst() float64
	GetDiff() float64
	GetPoints() int
	AddPoints(points int)
}

type TeamStats struct {
	ID            int
	Tournament    TournamentInterface
	Group         TournamentGroupInterface
	Team          TeamInterface
	Played        int
	Wins          int
	Losses        int
	Ties          int
	PointsFor     float64
	PointsAgainst float64
	Diff          float64
	Points        int
}

func (t *TeamStats) GetID() int {
	return t.ID
}

func (t *TeamStats) GetTournament() TournamentInterface {
	return t.Tournament
}

func (t *TeamStats) GetGroup() TournamentGroupInterface {
	return t.Group
}

func (t *TeamStats) GetTeam() TeamInterface {
	return t.Team
}

func (t *TeamStats) GetPlayed() int {
	return t.Played
}

func (t *TeamStats) GetWins() int {
	return t.Wins
}

func (t *TeamStats) GetLosses() int {
	return t.Losses
}

func (t *TeamStats) GetTies() int {
	return t.Ties
}

func (t *TeamStats) GetPointsFor() float64 {
	return t.PointsFor
}

func (t *TeamStats) GetPointsAgainst() float64 {
	return t.PointsAgainst
}

func (t *TeamStats) GetDiff() float64 {
	return t.Diff
}

func (t *TeamStats) GetPoints() int {
	return t.Points
}

func (t *TeamStats) AddPoints(points int) {
	t.Points += points
}
