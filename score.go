package gotournament

type ScoreInterface interface {
	GetID() int
	GetPoints() float64
	SetPoints(points float64)
}

type Score struct {
	ID     int
	Points float64 // We want to support any type of game where points can be very high or even just decimals
}

// 'id', 'game_id', 'team_id', 'score'

func (s Score) GetID() int {
	return s.ID
}

func (s Score) GetPoints() float64 {
	return s.Points
}

func (s Score) SetScore(points float64) {
	s.Points = points
}
