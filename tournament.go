package gotournament

import "fmt"

// TournamentInterface defines the methods needed to handle tournaments.
type TournamentInterface interface {
	GetType() int
	GetTypeString() string
	GetTeams() []TeamInterface
	GetGroups() []TournamentGroupInterface
	GetGames() []GameInterface
	Print() string
}

// Tournament is a default struct used as an example of how structs can be implemented for gotournament
type Tournament struct {
	Type   TournamentType // Is it elimination or group or ladder or poker? What is a type?
	Teams  []TeamInterface
	Groups []TournamentGroupInterface
	Games  []GameInterface
}

// GetType returns the type of tournament
func (t Tournament) GetType() int {
	return int(t.Type)
}

// GetType returns the type of tournament
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
func CreateTournament(teamCount int, meetCount int, tournamentType int) TournamentInterface {
	var teams []TeamInterface

	for i := 0; i < teamCount; i++ {
		teams = append(teams, &Team{ID: i})
	}

	return CreateTournamentFromTeams(teams, meetCount, tournamentType)
}

// CreateTournamentFromTeams takes a slice of teams and generates a tournament of the specified type
func CreateTournamentFromTeams(teams []TeamInterface, meetCount int, tournamentType int) TournamentInterface {
	if TournamentType(tournamentType) == TournamentTypeGroup {
		// It is recommended to call CreateGroupTournamentFromTeams directly as we try to automatically determine a group size and count here
		groupCount := len(teams) / 4 // We assume 4 teams per group by default
		if groupCount < 1 {
			groupCount = 1
		}
		return CreateGroupTournamentFromTeams(teams, groupCount, meetCount)
	} else if TournamentType(tournamentType) == TournamentTypeSeries {
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
	// We need to keep of eliminated teams, maybe make a function for that
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
		teams = append(teams, *group.GetTeams()...)
		gTeams := *group.GetTeams()
		// Loop through meet count
		for mi := 0; mi < meetCount; mi++ {
			// TODO a bug here causes team IDs to disappear with a meetcount of 4 or higher
			// TODO a bug here causes the same team to face itself with a meetcount of 4 or higher
			if len(gTeams) > 1 {
				var homeTeams []TeamInterface
				var awayTeams []TeamInterface
				// Everyone meets everyone once
				// We begin by taking our slice of teams like 0,1,2,3, and splitting it into home and away teams
				if len(gTeams) >= 4 {
					// if meet index is even
					if mi%2 == 0 {
						// The first half of the team slice become the home teams
						homeTeams = gTeams[0:(len(gTeams) / 2)]
						// The second half of the team slice become the away teams
						awayTeams = gTeams[(len(gTeams) / 2):]
						// if meet index is odd
					} else {
						awayTeams = gTeams[0:(len(gTeams) / 2)]
						homeTeams = gTeams[(len(gTeams) / 2):]
					}
				} else {
					var x TeamInterface
					// if meet index is even
					if mi%2 == 0 {
						// we take the team at index 0 and put the rest of the teams in the home team  slice
						x, homeTeams = gTeams[0], gTeams[1:]
						// The team that was first in the slice becomes the away team
						awayTeams = []TeamInterface{x}
						// if meet index is odd
					} else {
						x, awayTeams = gTeams[0], gTeams[1:]
						homeTeams = []TeamInterface{x}
					}
				}
				for i := 1; i < len(gTeams); i++ {
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
					if len(gTeams) >= 4 {
						homeTeams, awayTeams = RotateTeamsForCrossMatching(homeTeams, awayTeams, gTeams)
					} else {
						// We are dealing with less than 4 teams so we just switch sides
						tempTeams := homeTeams
						homeTeams = awayTeams
						awayTeams = tempTeams
					}
				}
			}
		}
	}
	return Tournament{Groups: groups, Games: games, Teams: teams, Type: TournamentTypeGroup}
}

func RotateTeamsForCrossMatching(homeTeams []TeamInterface, awayTeams []TeamInterface, gTeams []TeamInterface) ([]TeamInterface, []TeamInterface) {
	var x, y, z TeamInterface
	checkForDupe("g1", gTeams)
	// We keep the first home team in the same position and rotate all others
	// HT = Home Teams, AT = Away Teams
	// for HT 0,1 and AT 2,3. 0 is kept in place while 1 remains in the home team array
	x, homeTeams = homeTeams[0], homeTeams[1:]
	checkForDupe("g3", gTeams)
	// Take the first away team
	// 2 is taken out of AT, 3 remains in AT
	z, awayTeams = awayTeams[0], awayTeams[1:]
	checkForDupe("g3", gTeams)
	// and append to end of home teams
	// HT is now 1,2
	homeTeams = append(homeTeams, z)
	checkForDupe("g4", gTeams)
	// Take the first home team
	// 1 is taken out of HT, HT is now 2
	y, homeTeams = homeTeams[0], homeTeams[1:]
	checkForDupe("g5", gTeams)
	// and append it to the end of away teams
	// 1 is added to end of AT, AT is now 3,1
	fmt.Printf("g1 ")
	for _, t := range gTeams {
		fmt.Printf("%d ", t.GetID())
	}
	fmt.Println()
	fmt.Printf("away: ")
	for _, t := range awayTeams {
		fmt.Printf("%d ", t.GetID())
	}
	fmt.Println()
	fmt.Printf("home: ")
	for _, t := range homeTeams {
		fmt.Printf("%d ", t.GetID())
	}
	fmt.Println()
	fmt.Println("y", y.GetID())
	fmt.Println("x", x.GetID())
	awayTeams = append(awayTeams, y)
	fmt.Printf("g2 ")
	for _, t := range gTeams {
		fmt.Printf("%d ", t.GetID())
	}
	fmt.Println()
	fmt.Printf("away: ")
	for _, t := range awayTeams {
		fmt.Printf("%d ", t.GetID())
	}
	fmt.Println()
	fmt.Printf("home: ")
	for _, t := range homeTeams {
		fmt.Printf("%d ", t.GetID())
	}
	fmt.Println()
	checkForDupe("g6", gTeams)
	// Put the first home team back in first position of home array
	// HT is now 0,2
	homeTeams = append([]TeamInterface{x}, homeTeams...)
	checkForDupe("g7", gTeams)
	return homeTeams, awayTeams
}

// NumberOfGamesForGroupTournament Calculates the number of games in a group tournament based on number of teams, groups and unique encounters.
func NumberOfGamesForGroupTournament(teamCount int, groupCount int, meetCount int) int {
	return ((((teamCount / groupCount) - 1) * ((teamCount / groupCount) / 2)) * groupCount) * meetCount
}

func checkForDupe(key string, gTeams []TeamInterface) {
	var enc []int
	for _, t := range gTeams {
		for _, en := range enc {
			if en == t.GetID() {
				fmt.Printf("%s:  ", key)
				for _, t := range gTeams {
					fmt.Printf("%d ", t.GetID())
				}
				fmt.Println()
				panic("dupe")
			}
		}
		enc = append(enc, t.GetID())
	}
}
