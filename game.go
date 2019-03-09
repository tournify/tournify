package gotournament

type GameInterface interface {
	GetID() int
	GetTournament() TournamentInterface
	GetHomeTeam() TeamInterface // TODO home and away team could just be first or last team in team slice?
	GetAwayTeam() TeamInterface
	GetHomeScore() ScoreInterface
	GetAwayScore() ScoreInterface
	GetTeams() []TeamInterface   // For games that can have any number of teams
	GetScores() []ScoreInterface // For games that can have any number of scores
	SetScore(homeScore float64, awayScore float64)
}

type Game struct {
	ID         int
	Tournament TournamentInterface
	HomeTeam   TeamInterface
	AwayTeam   TeamInterface
	HomeScore  ScoreInterface
	AwayScore  ScoreInterface
	Scores     []ScoreInterface
	Teams      []TeamInterface
}

func (g Game) SetScore(homeScore float64, awayScore float64) {
	g.HomeScore.SetPoints(homeScore)
	g.AwayScore.SetPoints(awayScore)
}

func (g Game) GetID() int {
	return g.ID
}

func (g Game) GetTournament() TournamentInterface {
	return g.Tournament
}

func (g Game) GetHomeTeam() TeamInterface {
	return g.HomeTeam
}

func (g Game) GetAwayTeam() TeamInterface {
	return g.AwayTeam
}

func (g Game) GetHomeScore() ScoreInterface {
	return g.HomeScore
}

func (g Game) GetAwayScore() ScoreInterface {
	return g.AwayScore
}

func (g Game) GetTeams() []TeamInterface {
	return g.Teams
}

func (g Game) GetScores() []ScoreInterface {
	return g.Scores
}
