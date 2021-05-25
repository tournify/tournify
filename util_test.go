package tournify

import "testing"

func TestUtilEqualSlices(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}
	if !isEqual(s, s2) {
		t.Errorf("Slices are not equal %v != %v\n", s, s2)
	}
}

func TestUtilUnequalSlices(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	s2 := []int{5, 4, 3, 2, 1}
	if isEqual(s, s2) {
		t.Errorf("Slices are equal but expected not equal %v != %v\n", s, s2)
	}

	s3 := []int{1, 2, 3}
	s4 := []int{1, 2, 3, 4, 5}
	if isEqual(s3, s4) {
		t.Errorf("Slices are equal but expected not equal %v != %v\n", s, s2)
	}
}
