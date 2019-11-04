# GoTournament

[![GoDoc](https://godoc.org/github.com/tournify/gotournament?status.svg)](https://godoc.org/github.com/tournify/gotournament)
[![Go Report Card](https://goreportcard.com/badge/github.com/tournify/gotournament)](https://goreportcard.com/report/github.com/tournify/gotournament)
[![Build Status](https://api.travis-ci.org/tournify/gotournament.svg?branch=master)](https://travis-ci.org/tournify/gotournament)
[![Build status](https://ci.appveyor.com/api/projects/status/9s2ykpx3wdnf9eiw?svg=true)](https://ci.appveyor.com/project/markustenghamn/gotournament)
[![CircleCI](https://circleci.com/gh/tournify/gotournament.svg?style=svg)](https://circleci.com/gh/tournify/gotournament)
[![codecov](https://codecov.io/gh/tournify/gotournament/branch/master/graph/badge.svg)](https://codecov.io/gh/tournify/gotournament)

This project aims to support the creation of any tournament.

Current features
 - Group tournament creation
 - Group tournament stats
 
Planned features
 - Elimination tournaments
 - Double elimination
 - Round robin

Example
=

To create a group tournament with 2 groups where all teams in each group meet one time simply do the following.

```go
package main

import (
	"fmt"
	"github.com/tournify/gotournament"
)

func main()  {
	teams := []gotournament.Team{
		{ID:0},
		{ID:1},
		{ID:2},
		{ID:3},
		{ID:4},
		{ID:5},
		{ID:6},
		{ID:7},
		{ID:8},
		{ID:9},
		{ID:10},
		{ID:11},
		{ID:12},
		{ID:13},
		{ID:14},
		{ID:15},
	}

	teamInterfaces := make([]gotournament.TeamInterface, len(teams))

	for i := range teams {
		teamInterfaces[i] = &teams[i]
	}
    
    // The CreateGroupTournamentFromTeams method takes a slice of teams along with the group count and meet count
	tournament := gotournament.CreateGroupTournamentFromTeams(teamInterfaces, 2, 1)

    // The print method gives us a string representing the current tournament
	fmt.Println(tournament.Print())
}
```

This will print something similar to the following output.

```text
TournamentType 0

Groups
Group ID: 0
Team ID: 0
Team ID: 1
Team ID: 2
Team ID: 3
Team ID: 4
Team ID: 5
Team ID: 6
Team ID: 7

Group ID: 1
Team ID: 8
Team ID: 9
Team ID: 10
Team ID: 11
Team ID: 12
Team ID: 13
Team ID: 14
Team ID: 15


Games
Game ID: 0, HomeTeam: 0, AwayTeam: 4, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 1, AwayTeam: 5, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 2, AwayTeam: 6, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 3, AwayTeam: 7, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 0, AwayTeam: 5, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 2, AwayTeam: 6, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 3, AwayTeam: 7, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 4, AwayTeam: 1, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 0, AwayTeam: 6, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 3, AwayTeam: 7, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 4, AwayTeam: 1, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 5, AwayTeam: 2, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 0, AwayTeam: 7, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 4, AwayTeam: 1, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 5, AwayTeam: 2, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 6, AwayTeam: 3, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 0, AwayTeam: 1, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 5, AwayTeam: 2, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 6, AwayTeam: 3, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 7, AwayTeam: 4, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 0, AwayTeam: 2, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 6, AwayTeam: 3, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 7, AwayTeam: 4, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 1, AwayTeam: 5, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 0, AwayTeam: 3, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 7, AwayTeam: 4, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 1, AwayTeam: 5, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 2, AwayTeam: 6, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 8, AwayTeam: 12, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 9, AwayTeam: 13, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 10, AwayTeam: 14, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 11, AwayTeam: 15, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 8, AwayTeam: 13, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 10, AwayTeam: 14, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 11, AwayTeam: 15, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 12, AwayTeam: 9, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 8, AwayTeam: 14, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 11, AwayTeam: 15, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 12, AwayTeam: 9, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 13, AwayTeam: 10, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 8, AwayTeam: 15, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 12, AwayTeam: 9, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 13, AwayTeam: 10, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 14, AwayTeam: 11, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 8, AwayTeam: 9, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 13, AwayTeam: 10, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 14, AwayTeam: 11, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 15, AwayTeam: 12, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 8, AwayTeam: 10, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 14, AwayTeam: 11, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 15, AwayTeam: 12, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 9, AwayTeam: 13, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 8, AwayTeam: 11, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 15, AwayTeam: 12, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 9, AwayTeam: 13, HomeScore: 0.00, AwayScore: 0.00
Game ID: 0, HomeTeam: 10, AwayTeam: 14, HomeScore: 0.00, AwayScore: 0.00
```

Contributing
=

Please see [contributing](CONTRIBUTING.md).
