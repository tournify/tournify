package gotournament

type PlayerInterface interface {
	GetID() int
}

type Player struct {
	ID int
}

func (p Player) GetID() int {
	return p.ID
}
