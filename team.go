package gotournament

type TeamInterface interface {
	GetID() int
	GetPlayers() []PlayerInterface
	GetGames() []GameInterface
	AppendGame(game GameInterface)
}

type Team struct {
	ID      int
	Players []PlayerInterface
	Games   []GameInterface
}

func (t *Team) GetID() int {
	return t.ID
}

func (t *Team) GetPlayers() []PlayerInterface {
	return t.Players
}

func (t *Team) GetGames() []GameInterface {
	return t.Games
}

func (t *Team) AppendGame(game GameInterface) {
	t.Games = append(t.Games, game)
}
