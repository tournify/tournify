package gotournament

type Game struct {
	ID         int
	Tournament Tournament
	HomeTeam   Team
	AwayTeam   Team
	HomeScore  Score
	AwayScore  Score
	Scores     []Score // For games that can have more than any number of scores
}

func (g Game) SetScore(homeScore float64, awayScore float64) {
	g.HomeScore.Points = homeScore
	g.AwayScore.Points = awayScore
}
