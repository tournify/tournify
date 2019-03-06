package gotournament

import (
	"github.com/murlokswarm/errors"
	"sort"
)

type Tournament struct {
	ID     int
	Type   TournamentType // Is it elimination or group or ladder or poker? What is a type?
	Teams  []Team
	Groups []TournamentGroup
	Games  []Game
}

// TournamentGroup is for group tournaments only
type TournamentGroup struct {
	ID         int
	Tournament Tournament
	Teams      []Team
	Games      []Game
}

type TournamentTeamStats struct {
	ID            int
	Tournament    Tournament
	Group         TournamentGroup
	Team          Team
	Played        int
	Wins          int
	Losses        int
	Ties          int
	PointsFor     float64
	PointsAgainst float64
	Diff          float64
	Points        int
}

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

func CreateTournament(teamCount int, meetCount int, tournamentType TournamentType) Tournament {
	var teams []Team

	for i := 0; i < teamCount; i++ {
		teams = append(teams, Team{ID: i})
	}

	return CreateTournamentFromTeams(teams, meetCount, tournamentType)
}

func CreateTournamentFromTeams(teams []Team, meetCount int, tournamentType TournamentType) Tournament {
	if tournamentType == TournamentTypeGroup {
		// It is recommended to call CreateGroupTournamentFromTeams directly as we try to automatically determine a group size and count here
		groupCount := len(teams) / 4 // We assume 4 teams per group by default
		if groupCount < 1 {
			groupCount = 1
		}
		return CreateGroupTournamentFromTeams(teams, groupCount, meetCount)
	}
	return Tournament{}
}

func CreateGroupTournamentFromTeams(teams []Team, groupCount int, meetCount int) Tournament {
	var groups []TournamentGroup

	for i := 0; i < groupCount; i++ {
		groups = append(groups, TournamentGroup{ID: i})
	}

	for i, team := range teams {
		groupIndex := i % (len(groups))
		groups[groupIndex].Teams = append(groups[groupIndex].Teams, team)
	}

	return CreateGroupTournamentFromGroups(groups, meetCount)
}

func CreateGroupTournamentFromGroups(groups []TournamentGroup, meetCount int) Tournament {
	// Works best for an even amount of teams in every group
	var games []Game
	var teams []Team
	gameIndex := 0
	for gi, group := range groups {
		if len(group.Teams) < 4 {
			// TODO remove panic
			panic("Group must contain at least 4 teams, currently" + string(len(group.Teams)))
		}
		teams = append(teams, group.Teams...)
		// Loop through meet count
		for mi := 0; mi < meetCount; mi++ {
			var homeTeams []Team
			var awayTeams []Team
			// Everyone meets everyone once
			// We begin by taking our array of teams like 0,1,2,3, and splitting it
			if mi%2 == 0 {
				homeTeams = group.Teams[0:(len(group.Teams) / 2)]
				awayTeams = group.Teams[(len(group.Teams) / 2):]
			} else {
				awayTeams = group.Teams[0:(len(group.Teams) / 2)]
				homeTeams = group.Teams[(len(group.Teams) / 2):]
			}
			for i := 0; i < len(group.Teams)-1; i++ {
				// Now we have home teams of 0,1 and away teams of 2,3
				// This means 0 will meet 2 and 1 will meet 3
				for hi, hteam := range homeTeams {
					game := Game{HomeTeam: hteam, AwayTeam: awayTeams[hi]}
					groups[gi].Games = append(groups[gi].Games, game)
					hteam.Games = append(hteam.Games, game)
					games = append(games, game)
					awayTeams[hi].Games = append(awayTeams[hi].Games, game)
					gameIndex++
				}
				var x, y, z Team
				// We keep the first home team in the same position and rotate all others
				x, homeTeams = homeTeams[0], homeTeams[1:]
				// Take the first away team
				z, awayTeams = awayTeams[0], awayTeams[1:]
				// and append to end of home teams
				homeTeams = append(homeTeams, z)
				// Take the second home team
				y, homeTeams = homeTeams[0], homeTeams[1:]
				// and append it to the end of away teams
				awayTeams = append(awayTeams, y)
				// Put the first home team back in first position of home array
				homeTeams = append([]Team{x}, homeTeams...)
			}
		}
	}
	return Tournament{Groups: groups, Games: games, Teams: teams, Type: TournamentTypeGroup}
}

func NumberOfGames(teamCount int, groupCount int, meetCount int) int {
	return ((((teamCount / groupCount) - 1) * ((teamCount / groupCount) / 2)) * groupCount) * meetCount
}

func (t Tournament) GetGroupTournamentStats(winPoints int, lossPoints int, tiePoints int) (error, []TournamentTeamStats) {
	if t.Type != TournamentTypeGroup {
		return errors.New("Can not get stats for tournament type TournamentTypeGroup"), nil
	}
	var stats []TournamentTeamStats

	for _, group := range t.Groups {
		var groupStats []TournamentTeamStats

		for _, team := range group.Teams {
			stat := TournamentTeamStats{
				Group:         group,
				Team:          team,
				Played:        0,
				Wins:          0,
				Losses:        0,
				Ties:          0,
				PointsFor:     0.00,
				PointsAgainst: 0.00,
				Points:        0,
				Diff:          0.00}
			for _, game := range team.Games {
				if game.HomeTeam.ID == team.ID {
					stat.PointsFor = game.HomeScore.Points
					stat.PointsAgainst = game.AwayScore.Points
					if game.HomeScore.Points > game.AwayScore.Points {
						stat.Wins++
					} else if game.HomeScore.Points == game.AwayScore.Points {
						stat.Ties++
					} else {
						stat.Losses++
					}
				} else if game.AwayTeam.ID == team.ID {
					stat.PointsFor = game.AwayScore.Points
					stat.PointsAgainst = game.HomeScore.Points
					if game.HomeScore.Points < game.AwayScore.Points {
						stat.Wins++
					} else if game.HomeScore.Points == game.AwayScore.Points {
						stat.Ties++
					} else {
						stat.Losses++
					}
				}
				stat.Played++
			}
			stat.Points = stat.Wins * winPoints
			stat.Points += stat.Losses * lossPoints
			stat.Points += stat.Ties * tiePoints

			stat.Diff = stat.PointsFor - stat.PointsAgainst

			groupStats = append(groupStats, stat)
		}
		groupStats = SortTournamentStats(groupStats)
		stats = append(stats, groupStats...)
	}
	return nil, stats
}

func SortTournamentStats(stats []TournamentTeamStats) []TournamentTeamStats {
	sort.Slice(stats, func(i, j int) bool {
		if stats[i].Points > stats[j].Points {
			return true
		} else if stats[i].Points < stats[j].Points {
			return false
		} else {
			if stats[i].Diff > stats[j].Diff {
				return true
			} else if stats[i].Diff < stats[j].Diff {
				return false
			} else {
				if stats[i].PointsFor > stats[j].PointsFor {
					return true
				} else if stats[i].PointsFor < stats[j].PointsFor {
					return false
				}
			}
		}
		return true
	})
	return stats
}
