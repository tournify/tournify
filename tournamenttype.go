package gotournament

// TODO is this a good way of defining tournament types?
//  How can this be extended by others who want to build on this code without modifying the library?
type TournamentType int

const (
	TournamentTypeGroup             TournamentType = 0
	TournamentTypeElimination       TournamentType = 1
	TournamentTypeDoubleElimination TournamentType = 2
)

func (tournamentType TournamentType) String() string {
	names := [...]string{"Group", "Elimination", "DoubleElimination"}

	if tournamentType < TournamentTypeGroup || tournamentType > TournamentTypeDoubleElimination {
		return "Unknown"
	}

	return names[tournamentType]
}
