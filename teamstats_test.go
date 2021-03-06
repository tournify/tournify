package tournify

import (
	"testing"
)

func TestSortTournamentStats(t *testing.T) {
	var stats []TeamStatsInterface
	stats = append(stats, &TeamStats{Points: 10, PointsAgainst: 0, PointsFor: 10})
	stats = append(stats, &TeamStats{Points: 11, PointsAgainst: 0, PointsFor: 9})
	stats = append(stats, &TeamStats{Points: 10, PointsAgainst: 0, PointsFor: 8})
	stats = append(stats, &TeamStats{Points: 10, PointsAgainst: -2, PointsFor: 9})

	stats = SortTournamentStats(stats)

	for i, s := range stats {
		if i == 0 {
			verifyStats(s, 11, 9, 9, 0, t)
		} else if i == 1 {
			verifyStats(s, 10, 11, 9, -2, t)
		} else if i == 2 {
			verifyStats(s, 10, 10, 10, 0, t)
		}
	}

}

func TestGroupTournamentStats(t *testing.T) {
	teamCount := 8
	meetCount := 2
	groupCount := 2
	tournament := CreateTournament(teamCount, meetCount, groupCount, int(TournamentTypeGroup))

	for _, group := range tournament.GetGroups() {
		gGames := *group.GetGames()
		homeScore := 0
		awayScore := len(gGames)
		for i := range gGames {
			gGames[i].SetScore(float64(homeScore), float64(awayScore))
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
			verifyStats(s, 13, 28, 50, 22, t)
		} else if i == 1 {
			verifyStats(s, 10, -4, 34, 38, t)
		} else if i == 2 {
			verifyStats(s, 9, -4, 34, 38, t)
		} else if i == 3 {
			verifyStats(s, 3, -20, 26, 46, t)
		} else if i == 4 {
			verifyStats(s, 13, 28, 50, 22, t)
		} else if i == 5 {
			verifyStats(s, 10, -4, 34, 38, t)
		} else if i == 6 {
			verifyStats(s, 9, -4, 34, 38, t)
		} else if i == 7 {
			verifyStats(s, 3, -20, 26, 46, t)
		}
	}

}

func verifyStats(stats TeamStatsInterface, points int, diff float64, pointsfor float64, pointsAgainst float64, t *testing.T) {
	if stats.GetPoints() != points {
		t.Errorf("Expected points %d, currently: %d", points, stats.GetPoints())
	}
	if stats.GetDiff() != diff {
		t.Errorf("Expected diff %f, currently: %f", diff, stats.GetDiff())
	}
	if stats.GetPointsFor() != pointsfor {
		t.Errorf("Expected points for to be %f, currently: %f", pointsfor, stats.GetPointsFor())
	}
	if stats.GetPointsAgainst() != pointsAgainst {
		t.Errorf("Expected points against to be %f, currently: %f", pointsAgainst, stats.GetPointsAgainst())
	}
}
