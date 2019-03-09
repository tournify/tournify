package gotournament

import (
	"github.com/murlokswarm/errors"
	"sort"
)

type TournamentInterface interface {
	GetID() int
	GetType() int
	GetTeams() []TeamInterface
	GetGroups() []TournamentGroupInterface
	GetGames() []GameInterface
}

type Tournament struct {
	ID     int
	Type   TournamentType // Is it elimination or group or ladder or poker? What is a type?
	Teams  []TeamInterface
	Groups []TournamentGroupInterface
	Games  []GameInterface
}

func (t Tournament) GetID() int {
	return t.ID
}

func (t Tournament) GetType() int {
	return int(t.Type)
}

func (t Tournament) GetTeams() []TeamInterface {
	return t.Teams
}

func (t Tournament) GetGroups() []TournamentGroupInterface {
	return t.Groups
}

func (t Tournament) GetGames() []GameInterface {
	return t.Games
}

func CreateTournament(teamCount int, meetCount int, tournamentType int) TournamentInterface {
	var teams []TeamInterface

	for i := 0; i < teamCount; i++ {
		teams = append(teams, &Team{ID: i})
	}

	return CreateTournamentFromTeams(teams, meetCount, tournamentType)
}

func CreateTournamentFromTeams(teams []TeamInterface, meetCount int, tournamentType int) TournamentInterface {
	if TournamentType(tournamentType) == TournamentTypeGroup {
		// It is recommended to call CreateGroupTournamentFromTeams directly as we try to automatically determine a group size and count here
		groupCount := len(teams) / 4 // We assume 4 teams per group by default
		if groupCount < 1 {
			groupCount = 1
		}
		return CreateGroupTournamentFromTeams(teams, groupCount, meetCount)
	}
	return nil
}

func CreateGroupTournamentFromTeams(teams []TeamInterface, groupCount int, meetCount int) TournamentInterface {
	var groups []TournamentGroupInterface

	for i := 0; i < groupCount; i++ {
		groups = append(groups, &TournamentGroup{ID: i})
	}

	for i, team := range teams {
		groupIndex := i % (len(groups))
		groups[groupIndex].AppendTeam(team)
	}

	return CreateGroupTournamentFromGroups(groups, meetCount)
}

func CreateGroupTournamentFromGroups(groups []TournamentGroupInterface, meetCount int) TournamentInterface {
	// Works best for an even amount of teams in every group
	var games []GameInterface
	var teams []TeamInterface
	gameIndex := 0
	for gi, group := range groups {
		groupteams := group.GetTeams()
		if len(groupteams) < 4 {
			// TODO remove panic
			panic("Group must contain at least 4 teams, currently" + string(len(groupteams)))
		}
		teams = append(teams, group.GetTeams()...)
		// Loop through meet count
		for mi := 0; mi < meetCount; mi++ {
			var homeTeams []TeamInterface
			var awayTeams []TeamInterface
			// Everyone meets everyone once
			// We begin by taking our array of teams like 0,1,2,3, and splitting it
			if mi%2 == 0 {
				homeTeams = group.GetTeams()[0:(len(group.GetTeams()) / 2)]
				awayTeams = group.GetTeams()[(len(group.GetTeams()) / 2):]
			} else {
				awayTeams = group.GetTeams()[0:(len(group.GetTeams()) / 2)]
				homeTeams = group.GetTeams()[(len(group.GetTeams()) / 2):]
			}
			for i := 0; i < len(group.GetTeams())-1; i++ {
				// Now we have home teams of 0,1 and away teams of 2,3
				// This means 0 will meet 2 and 1 will meet 3
				for hi, hteam := range homeTeams {
					game := Game{HomeTeam: hteam, AwayTeam: awayTeams[hi]}
					groups[gi].AppendGame(game)
					hteam.AppendGame(game)
					games = append(games, game)
					awayTeams[hi].AppendGame(game)
					gameIndex++
				}
				var x, y, z TeamInterface
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
				homeTeams = append([]TeamInterface{x}, homeTeams...)
			}
		}
	}
	return Tournament{Groups: groups, Games: games, Teams: teams, Type: TournamentTypeGroup}
}

func NumberOfGames(teamCount int, groupCount int, meetCount int) int {
	return ((((teamCount / groupCount) - 1) * ((teamCount / groupCount) / 2)) * groupCount) * meetCount
}

func (t Tournament) GetGroupTournamentStats(winPoints int, lossPoints int, tiePoints int) (error, []TeamStatsInterface) {
	if t.Type != TournamentTypeGroup {
		return errors.New("Can not get stats for tournament type TournamentTypeGroup"), nil
	}
	var stats []TeamStatsInterface

	for _, group := range t.Groups {
		var groupStats []TeamStatsInterface

		for _, team := range group.GetTeams() {
			stat := TeamStats{
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
			for _, game := range team.GetGames() {
				if game.GetHomeTeam().GetID() == team.GetID() {
					stat.PointsFor = game.GetHomeScore().GetPoints()
					stat.PointsAgainst = game.GetAwayScore().GetPoints()
					if game.GetHomeScore().GetPoints() > game.GetAwayScore().GetPoints() {
						stat.Wins++
					} else if game.GetHomeScore().GetPoints() == game.GetAwayScore().GetPoints() {
						stat.Ties++
					} else {
						stat.Losses++
					}
				} else if game.GetAwayTeam().GetID() == team.GetID() {
					stat.PointsFor = game.GetAwayScore().GetPoints()
					stat.PointsAgainst = game.GetHomeScore().GetPoints()
					if game.GetHomeScore().GetPoints() < game.GetAwayScore().GetPoints() {
						stat.Wins++
					} else if game.GetHomeScore().GetPoints() == game.GetAwayScore().GetPoints() {
						stat.Ties++
					} else {
						stat.Losses++
					}
				}
				stat.Played++
			}
			stat.AddPoints(stat.Wins * winPoints)
			stat.AddPoints(stat.Losses * lossPoints)
			stat.AddPoints(stat.Ties * tiePoints)

			stat.Diff = stat.GetPointsFor() - stat.GetPointsAgainst()

			groupStats = append(groupStats, &stat)
		}
		groupStats = SortTournamentStats(groupStats)
		stats = append(stats, groupStats...)
	}
	return nil, stats
}

func SortTournamentStats(stats []TeamStatsInterface) []TeamStatsInterface {
	sort.Slice(stats, func(i, j int) bool {
		if stats[i].GetPoints() > stats[j].GetPoints() {
			return true
		} else if stats[i].GetPoints() < stats[j].GetPoints() {
			return false
		} else {
			if stats[i].GetDiff() > stats[j].GetDiff() {
				return true
			} else if stats[i].GetDiff() < stats[j].GetDiff() {
				return false
			} else {
				if stats[i].GetPointsFor() > stats[j].GetPointsFor() {
					return true
				} else if stats[i].GetPointsFor() < stats[j].GetPointsFor() {
					return false
				}
			}
		}
		return true
	})
	return stats
}
