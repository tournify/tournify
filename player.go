package gotournament

// PlayerInterface defines the needed methods for players used by this library
type PlayerInterface interface {
	GetID() int
}

// Player is a default struct used as an example of how structs can be implemented for gotournament
type Player struct {
	ID int
}

// GetID returns the id of the player
func (p Player) GetID() int {
	return p.ID
}
