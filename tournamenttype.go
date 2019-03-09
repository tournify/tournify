package gotournament

// TournamentType defines the type of tournament
// TODO is this a good way of defining tournament types?
//  How can this be extended by others who want to build on this code without modifying the library?
type TournamentType int

const (
	// TournamentTypeGroup is for group tournaments
	TournamentTypeGroup TournamentType = 0
	// TournamentTypeElimination is for elimination or knockout tournaments
	TournamentTypeElimination TournamentType = 1
	// TournamentTypeDoubleElimination is the same as TournamentTypeElimination
	// but teams to get knocked out early get a second chance to come back and win
	TournamentTypeDoubleElimination TournamentType = 2
)

func (tournamentType TournamentType) String() string {
	names := [...]string{"Group", "Elimination", "DoubleElimination"}

	if tournamentType < TournamentTypeGroup || tournamentType > TournamentTypeDoubleElimination {
		return "Unknown"
	}

	return names[tournamentType]
}
