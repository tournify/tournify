package gotournament

import (
	"errors"
	"sort"
)

// TeamStatsInterface is used to show team statistics. Currently this is specifically made for
// group tournaments where there is a need to rank teams.
type TeamStatsInterface interface {
	GetGroup() TournamentGroupInterface
	GetTeam() TeamInterface
	GetPlayed() int
	GetWins() int
	GetLosses() int
	GetTies() int
	GetPointsFor() float64
	GetPointsAgainst() float64
	GetDiff() float64
	GetPoints() int
	AddPoints(points int)
}

// TeamStats is a default struct used as an example of how structs can be implemented for gotournament
type TeamStats struct {
	Tournament    TournamentInterface
	Group         TournamentGroupInterface
	Team          TeamInterface
	Played        int
	Wins          int
	Losses        int
	Ties          int
	PointsFor     float64
	PointsAgainst float64
	Points        int
}

// GetGroup returns the Group that the statistics were generated for, stats are directly related to a team and the group they are in.
func (t *TeamStats) GetGroup() TournamentGroupInterface {
	return t.Group
}

// GetTeam returns the Team that the statistics were generated for, stats are directly related to a team and the group they are in.
func (t *TeamStats) GetTeam() TeamInterface {
	return t.Team
}

// GetPlayed returns the number of games played
func (t *TeamStats) GetPlayed() int {
	return t.Played
}

// GetWins returns the number of won games
func (t *TeamStats) GetWins() int {
	return t.Wins
}

// GetLosses returns the number of lost games
func (t *TeamStats) GetLosses() int {
	return t.Losses
}

// GetTies returns the number of games resulting in a tied game
func (t *TeamStats) GetTies() int {
	return t.Ties
}

// GetPointsFor returns the number of goals or points that this team has made
func (t *TeamStats) GetPointsFor() float64 {
	return t.PointsFor
}

// GetPointsAgainst returns the number of goals or points that other teams have made against this team
func (t *TeamStats) GetPointsAgainst() float64 {
	return t.PointsAgainst
}

// GetDiff returns the difference of PointsFor and PointsAgainst
func (t *TeamStats) GetDiff() float64 {
	return t.PointsFor - t.PointsAgainst
}

// GetPoints returns the number of points the team has based on wins, losses or ties
func (t *TeamStats) GetPoints() int {
	return t.Points
}

// AddPoints adds the specified number of points to Points
func (t *TeamStats) AddPoints(points int) {
	t.Points += points
}

// GetGroupTournamentStats takes 4 inouts. The first input is the tournament itself.
// The other three input defines how many points a team should get for a win, loss or tie. The standard is 3, 0, 1 but
// it can vary depending on the tournament.
func GetGroupTournamentStats(t TournamentInterface, winPoints int, lossPoints int, tiePoints int) ([]TeamStatsInterface, error) {
	if t.GetType() != int(TournamentTypeGroup) {
		return nil, errors.New("can not get stats for tournament type TournamentTypeGroup")
	}
	var stats []TeamStatsInterface

	for _, group := range t.GetGroups() {
		var groupStats []TeamStatsInterface

		for _, team := range *group.GetTeams() {
			stat := TeamStats{
				Group:         group,
				Team:          team,
				Played:        0,
				Wins:          0,
				Losses:        0,
				Ties:          0,
				PointsFor:     0.00,
				PointsAgainst: 0.00,
				Points:        0}
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

			groupStats = append(groupStats, &stat)
		}
		groupStats = sortTournamentStats(groupStats)
		stats = append(stats, groupStats...)
	}
	return stats, nil
}

func sortTournamentStats(stats []TeamStatsInterface) []TeamStatsInterface {
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
