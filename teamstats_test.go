package gotournament

import (
	"testing"
)

func TestSortTournamentStats(t *testing.T) {
	var stats []TeamStatsInterface
	stats = append(stats, &TeamStats{Points: 10, PointsAgainst: 0, PointsFor: 10})
	stats = append(stats, &TeamStats{Points: 11, PointsAgainst: 0, PointsFor: 9})
	stats = append(stats, &TeamStats{Points: 10, PointsAgainst: 0, PointsFor: 8})
	stats = append(stats, &TeamStats{Points: 10, PointsAgainst: -2, PointsFor: 9})

	stats = sortTournamentStats(stats)

	for i, s := range stats {
		if i == 0 {
			verifyStats(s, 11, 9, 9, t)
		} else if i == 1 {
			verifyStats(s, 10, 11, 9, t)
		} else if i == 2 {
			verifyStats(s, 10, 10, 10, t)
		}
	}

}

func TestGroupTournamentStats(t *testing.T) {
	teamCount := 8
	meetCount := 2
	tournament := CreateTournament(teamCount, meetCount, int(TournamentTypeGroup))

	for _, group := range tournament.GetGroups() {
		homeScore := 0
		awayScore := len(group.GetGames())
		for i := range group.GetGames() {
			group.GetGames()[i].SetScore(float64(homeScore), float64(awayScore))
			homeScore++
			awayScore--
		}
	}

	stats, err := GetGroupTournamentStats(tournament, 3, 0, 1)

	if err != nil {
		t.Error(err)
	}

	for i, s := range stats {
		if i == 0 {
			verifyStats(s, 9, 10, 11, t)
		} else if i == 1 {
			verifyStats(s, 9, -8, 2, t)
		} else if i == 2 {
			verifyStats(s, 4, -10, 1, t)
		} else if i == 3 {
			verifyStats(s, 4, -10, 1, t)
		} else if i == 4 {
			verifyStats(s, 9, 10, 11, t)
		} else if i == 5 {
			verifyStats(s, 9, -8, 2, t)
		} else if i == 6 {
			verifyStats(s, 4, -10, 1, t)
		} else if i == 7 {
			verifyStats(s, 4, -10, 1, t)
		}
	}

}

func verifyStats(stats TeamStatsInterface, points int, diff float64, pointsfor float64, t *testing.T) {
	if stats.GetPoints() != points {
		t.Errorf("Expected points %d, currently: %d", points, stats.GetPoints())
	}
	if stats.GetDiff() != diff {
		t.Errorf("Expected diff %f, currently: %f", diff, stats.GetDiff())
	}
	if stats.GetPointsFor() != pointsfor {
		t.Errorf("Expected points %f, currently: %f", pointsfor, stats.GetPointsFor())
	}
}
