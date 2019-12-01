package tournify

import (
	"fmt"
	"math"
)

// TournamentInterface defines the methods needed to handle tournaments.
type TournamentInterface interface {
	GetType() int
	GetTypeString() string
	GetTeams() []TeamInterface
	GetEliminatedTeams() []TeamInterface // For elimination style tournaments
	GetRemainingTeams() []TeamInterface  // For elimination style tournaments
	GetGroups() []TournamentGroupInterface
	GetGames() []GameInterface
	Print() string
}

// Tournament is a default struct used as an example of how structs can be implemented for tournify
type Tournament struct {
	Type   TournamentType // Is it elimination or group or ladder or poker? What is a type?
	Teams  []TeamInterface
	Groups []TournamentGroupInterface
	Games  []GameInterface
}

// GetType returns the type of tournament as an int
func (t Tournament) GetType() int {
	return int(t.Type)
}

// GetTypeString returns the type of tournament as a string
func (t Tournament) GetTypeString() string {
	return t.Type.String()
}

// GetTeams returns the team slice
func (t Tournament) GetTeams() []TeamInterface {
	return t.Teams
}

// GetGroups returns the group slice
func (t Tournament) GetGroups() []TournamentGroupInterface {
	return t.Groups
}

// GetGames returns the game slice
func (t Tournament) GetGames() []GameInterface {
	return t.Games
}

// GetEliminatedTeams gets all teams that have been eliminated at least one time in an elimination tournament
func (t Tournament) GetEliminatedTeams() []TeamInterface {
	var elimnatedTeams []TeamInterface
	for _, team := range t.GetTeams() {
		if team.GetEliminatedCount() > 0 && t.GetType() == int(TournamentTypeElimination) {
			elimnatedTeams = append(elimnatedTeams, team)
		}
		if team.GetEliminatedCount() > 1 && t.GetType() == int(TournamentTypeDoubleElimination) {
			elimnatedTeams = append(elimnatedTeams, team)
		}
	}
	return elimnatedTeams
}

// GetRemainingTeams gets all teams that have not been eliminated in an elimination tournament
func (t Tournament) GetRemainingTeams() []TeamInterface {
	var remainingTeams []TeamInterface
	for _, team := range t.GetTeams() {
		if team.GetEliminatedCount() < 1 && t.GetType() == int(TournamentTypeElimination) {
			remainingTeams = append(remainingTeams, team)
		}
		if team.GetEliminatedCount() < 2 && t.GetType() == int(TournamentTypeDoubleElimination) {
			remainingTeams = append(remainingTeams, team)
		}
	}
	return remainingTeams
}

// Print writes the full tournament details to a string
func (t Tournament) Print() string {
	p := fmt.Sprintf("TournamentType: %s\n", t.GetTypeString())
	if t.GetType() == 0 {
		p += fmt.Sprintf("\nGroups\n")
		for _, group := range t.GetGroups() {
			p += group.Print()
		}
	} else {
		p += fmt.Sprintf("\nTeams\n")
		for _, team := range t.GetTeams() {
			p += team.Print()
		}
	}
	p += fmt.Sprintf("\nGames\n")
	for _, games := range t.GetGames() {
		p += games.Print()
	}
	return p
}

// CreateTournament creates a tournament with the simplest input. It is recommended to create a slice with
// specific use via CreateTournamentFromTeams as this method will generate it's own Teams as a sort of placeholder.
func CreateTournament(teamCount int, meetCount int, groupCount int, tournamentType int) TournamentInterface {
	var teams []TeamInterface

	for i := 0; i < teamCount; i++ {
		teams = append(teams, &Team{ID: i})
	}

	return CreateTournamentFromTeams(teams, meetCount, groupCount, tournamentType)
}

// CreateTournamentFromTeams takes a slice of teams and generates a tournament of the specified type
func CreateTournamentFromTeams(teams []TeamInterface, meetCount int, groupCount int, tournamentType int) TournamentInterface {
	if TournamentType(tournamentType) == TournamentTypeGroup {
		if groupCount < 1 {
			return nil
		}
		if meetCount < 1 {
			return nil
		}
		return CreateGroupTournamentFromTeams(teams, groupCount, meetCount)
	} else if TournamentType(tournamentType) == TournamentTypeSeries {
		// TODO this should return an tournament of type series
		return CreateGroupTournamentFromTeams(teams, 1, meetCount)
	} else if TournamentType(tournamentType) == TournamentTypeElimination {
		return CreateEliminationTournamentFromTeams(teams)
	}
	return nil
}

// CreateEliminationTournamentFromTeams takes a slice of teams and generates a elimination tournament
func CreateEliminationTournamentFromTeams(teams []TeamInterface) TournamentInterface {
	// Create the initial games of the elimination tournament
	var games []GameInterface
	// We need to keep track of eliminated teams, maybe make a function for that
	// also a function for teams still in the tournament
	// A function to calculate which team proceeds as well and generate the next game
	// Return a tournament
	return Tournament{Games: games, Teams: teams, Type: TournamentTypeElimination}
}

// CreateGroupTournamentFromTeams takes a slice of teams and generates a group tournament
func CreateGroupTournamentFromTeams(teams []TeamInterface, groupCount int, meetCount int) TournamentInterface {
	// TODO implement better error handling
	if groupCount < 1 || meetCount < 1 {
		return nil
	}

	var groups []TournamentGroupInterface
	teamsPerGroup := len(teams) / groupCount

	for i := 0; i < groupCount; i++ {
		groups = append(groups, &TournamentGroup{ID: i})
	}

	groupIndex := 0
	for i, team := range teams {
		adjGI := groupIndex + 1
		if i >= teamsPerGroup*adjGI && adjGI < groupCount {
			groupIndex++
		}
		groups[groupIndex].AppendTeam(team)
	}

	return CreateGroupTournamentFromGroups(groups, meetCount)
}

// CreateGroupTournamentFromGroups takes a slice of groups that contain teams and returns a group tournament
// TODO simplify and break down this function in to smaller chunks?
// TODO this method currently uses cross matching for games but other types of matching could be supported
func CreateGroupTournamentFromGroups(groups []TournamentGroupInterface, meetCount int) TournamentInterface {
	// Works best for an even amount of teams in every group
	var games []GameInterface
	var teams []TeamInterface
	gameIndex := 0
	for gi, group := range groups {
		var tempID int
		uneven := false

		teams = append(teams, *group.GetTeams()...)
		gTeams := *group.GetTeams()

		// If there is an uneven amount of teams we need to add a temporary team which is later removed
		if len(gTeams)%2 != 0 {
			tempID = generateTempID(gTeams, -1)
			tempTeam := Team{ID: tempID}
			gTeams = append(gTeams, &tempTeam)
			uneven = true
		}

		// Loop through meet count
		for mi := 0; mi < meetCount; mi++ {
			// TODO game calculation is wrong when there is an uneven number of teams per group
			if len(gTeams) > 1 {
				halfCountHiger := DivideRoundUp(len(gTeams), 2)
				halfCountLower := DivideRoundDown(len(gTeams), 2)
				homeTeams := make([]TeamInterface, halfCountHiger)
				awayTeams := make([]TeamInterface, halfCountLower)
				// Everyone meets everyone once
				// We begin by taking our slice of teams like 0,1,2,3, and splitting it into home and away teams
				// if meet index is even
				if mi%2 == 0 {
					// The first half of the team slice become the home teams
					copy(homeTeams, gTeams[0:halfCountHiger])
					// The second half of the team slice become the away teams
					copy(awayTeams, gTeams[halfCountHiger:])
					// if meet index is odd
				} else {
					copy(awayTeams, gTeams[0:halfCountHiger])
					copy(homeTeams, gTeams[halfCountLower:])
				}

				awayTeams = reverseSlice(awayTeams)

				for i := 0; i < len(gTeams)-1; i++ {
					// Now we have home teams of 0,1 and away teams of 2,3
					// This means 0 will meet 2 and 1 will meet 3
					for hi, hteam := range homeTeams {
						game := Game{Teams: []TeamInterface{hteam, awayTeams[hi]}}
						groups[gi].AppendGame(&game)
						hteam.AppendGame(&game)
						games = append(games, &game)
						awayTeams[hi].AppendGame(&game)
						gameIndex++
					}
					homeTeams, awayTeams = rotateTeamsForCrossMatching(homeTeams, awayTeams)

				}
			}
		}
		if uneven {
			games = removeTempGames(games, tempID)
		}
	}
	return Tournament{Groups: groups, Games: games, Teams: teams, Type: TournamentTypeGroup}
}

func removeTempGames(games []GameInterface, tempID int) []GameInterface {
	for i := 0; i < len(games); i++ {
		if games[i].GetHomeTeam().GetID() == tempID || games[i].GetAwayTeam().GetID() == tempID {
			return removeTempGames(append(games[:i], games[i+1:]...), tempID)
		}
	}
	return games
}

func generateTempID(teams []TeamInterface, tempID int) int {
	for _, t := range teams {
		if t.GetID() == tempID {
			return generateTempID(teams, tempID-1)
		}
	}
	return tempID
}

func reverseSlice(a []TeamInterface) []TeamInterface {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}

func rotateTeamsForCrossMatching(homeTeams []TeamInterface, awayTeams []TeamInterface) ([]TeamInterface, []TeamInterface) {
	var x, y, z TeamInterface
	// We keep the first home team in the same position and rotate all others
	// HT = Home Teams, AT = Away Teams
	// for HT 0,1 and AT 2,3. 0 is kept in place while 1 remains in the home team array
	x, homeTeams = homeTeams[0], homeTeams[1:]
	// Take the first away team
	// 2 is taken out of AT, 3 remains in AT
	z, awayTeams = awayTeams[0], awayTeams[1:]
	// and append to end of home teams
	// HT is now 1,2
	homeTeams = append(homeTeams, z)
	// Take the first home team
	// 1 is taken out of HT, HT is now 2
	y, homeTeams = homeTeams[0], homeTeams[1:]
	// and append it to the end of away teams
	// 1 is added to end of AT, AT is now 3,1
	awayTeams = append(awayTeams, y)
	// Put the first home team back in first position of home array
	// HT is now 0,2
	homeTeams = append([]TeamInterface{x}, homeTeams...)
	return homeTeams, awayTeams
}

// NumberOfGamesForGroupTournament Calculates the number of games in a group tournament based on number of teams, groups and unique encounters.
func NumberOfGamesForGroupTournament(teamCount int, groupCount int, meetCount int) int {
	tpg := float64(teamCount) / float64(groupCount)
	games := tpg * (tpg - 1) / 2
	res := int(games * float64(meetCount*groupCount))
	if math.Mod(float64(teamCount), float64(groupCount)) != 0 {
		res += int(math.Mod(float64(teamCount), float64(groupCount))) * meetCount
	}
	return res
}

// DivideRoundUp takes two ints, divides them and rounds the result up to the nearest int
func DivideRoundUp(a int, b int) int {
	return int(math.Ceil(float64(a) / float64(b)))
}

// DivideRoundDown takes two ints, divides them and rounds the result up to the nearest int
func DivideRoundDown(a int, b int) int {
	return int(math.Floor(float64(a) / float64(b)))
}
