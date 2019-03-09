package gotournament

import (
	"errors"
	"sort"
)

type TeamStatsInterface interface {
	GetTournament() TournamentInterface
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
	Diff          float64
	Points        int
}

func (t *TeamStats) GetTournament() TournamentInterface {
	return t.Tournament
}

func (t *TeamStats) GetGroup() TournamentGroupInterface {
	return t.Group
}

func (t *TeamStats) GetTeam() TeamInterface {
	return t.Team
}

func (t *TeamStats) GetPlayed() int {
	return t.Played
}

func (t *TeamStats) GetWins() int {
	return t.Wins
}

func (t *TeamStats) GetLosses() int {
	return t.Losses
}

func (t *TeamStats) GetTies() int {
	return t.Ties
}

func (t *TeamStats) GetPointsFor() float64 {
	return t.PointsFor
}

func (t *TeamStats) GetPointsAgainst() float64 {
	return t.PointsAgainst
}

func (t *TeamStats) GetDiff() float64 {
	return t.Diff
}

func (t *TeamStats) GetPoints() int {
	return t.Points
}

func (t *TeamStats) AddPoints(points int) {
	t.Points += points
}

func GetGroupTournamentStats(t TournamentInterface, winPoints int, lossPoints int, tiePoints int) (error, []TeamStatsInterface) {
	if t.GetType() != int(TournamentTypeGroup) {
		return errors.New("can not get stats for tournament type TournamentTypeGroup"), nil
	}
	var stats []TeamStatsInterface

	for _, group := range t.GetGroups() {
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
		groupStats = sortTournamentStats(groupStats)
		stats = append(stats, groupStats...)
	}
	return nil, stats
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
