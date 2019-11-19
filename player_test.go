package gotournament

import "testing"

func TestPlayerGetID(t *testing.T) {
	id := 42
	player := Player{ID: id}
	if player.GetID() != id {
		t.Errorf("ID does not match %d != %d\n", player.GetID(), id)
	}
}
