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
	Scores     []ScoreInterface
	Teams      []TeamInterface
}

func (g *Game) SetScore(homeScore float64, awayScore float64) {
	if len(g.Scores) < 1 {
		g.Scores = append(g.Scores, &Score{}, &Score{})
	} else if len(g.Scores) < 2 {
		g.Scores = append(g.Scores, &Score{})
	}
	g.Scores[0].SetPoints(homeScore)
	g.Scores[1].SetPoints(awayScore)
}

func (g *Game) GetID() int {
	return g.ID
}

func (g *Game) GetTournament() TournamentInterface {
	return g.Tournament
}

func (g *Game) GetHomeTeam() TeamInterface {
	if len(g.Scores) < 1 {
		g.Teams = append(g.Teams, &Team{})
	}
	return g.Teams[0]
}

func (g *Game) GetAwayTeam() TeamInterface {
	if len(g.Scores) < 1 {
		g.Teams = append(g.Teams, &Team{}, &Team{})
	} else if len(g.Scores) < 2 {
		g.Teams = append(g.Teams, &Team{})
	}
	return g.Teams[1]
}

func (g *Game) GetHomeScore() ScoreInterface {
	if len(g.Scores) < 1 {
		g.Scores = append(g.Scores, &Score{})
	}
	return g.Scores[0]
}

func (g *Game) GetAwayScore() ScoreInterface {
	if len(g.Scores) < 1 {
		g.Scores = append(g.Scores, &Score{}, &Score{})
	} else if len(g.Scores) < 2 {
		g.Scores = append(g.Scores, &Score{})
	}
	return g.Scores[1]
}

func (g *Game) GetTeams() []TeamInterface {
	return g.Teams
}

func (g *Game) GetScores() []ScoreInterface {
	return g.Scores
}
