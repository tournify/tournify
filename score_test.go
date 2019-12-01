package tournify

import (
	"testing"
)

func TestScoreGetID(t *testing.T) {
	id := 42
	score := Score{ID: id}
	if score.GetID() != id {
		t.Errorf("ID does not match %d != %d\n", score.GetID(), id)
	}
}

func TestScoreGetPoints(t *testing.T) {
	points := 10.0
	score := Score{ID: 1, Points: points}
	if score.GetPoints() != points {
		t.Errorf("Points do not match %.2f != %.2f\n", score.GetPoints(), points)
	}
}

func TestScoreSetPoints(t *testing.T) {
	points := 10.0
	score := Score{ID: 1}
	score.SetPoints(points)
	if score.GetPoints() != points {
		t.Errorf("Points do not match %.2f != %.2f\n", score.GetPoints(), points)
	}
}
