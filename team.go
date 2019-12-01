package tournify

import "fmt"

// TeamInterface defines the methods for teams. Teams are used to create tournaments and generate games.
// Teams can have games and can contain a slice of players
type TeamInterface interface {
	GetID() int
	GetPlayers() []PlayerInterface
	GetGames() []GameInterface
	GetEliminatedCount() int
	AppendGame(game GameInterface)
	Print() string
}

// Team is a default struct used as an example of how structs can be implemented for tournify
type Team struct {
	ID         int
	Players    []PlayerInterface
	Games      []GameInterface
	Eliminated int // Increment by 1 every time this team is elimnated
}

// GetID returns the id of the score
func (t *Team) GetID() int {
	return t.ID
}

// GetPlayers returns a slice of players
func (t *Team) GetPlayers() []PlayerInterface {
	return t.Players
}

// GetGames returns a slice of games
func (t *Team) GetGames() []GameInterface {
	return t.Games
}

// AppendGame takes a game as an argument and appends it to the Games slice
func (t *Team) AppendGame(game GameInterface) {
	t.Games = append(t.Games, game)
}

// Print writes team details to stdout
func (t *Team) Print() string {
	return fmt.Sprintf("Team ID: %d\n", t.GetID())
}

// GetEliminatedCount gets the number of times the team has been eliminated in a multiple elimination tournament
func (t *Team) GetEliminatedCount() int {
	return t.Eliminated
}
