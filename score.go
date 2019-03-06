package gotournament

type Score struct {
	ID     int
	Points float64 // We want to support any type of game where points can be very high or even just decimals
}

// 'id', 'game_id', 'team_id', 'score'
