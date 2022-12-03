package Set

import "testing"

func TestSetAddInt(t *testing.T) {
	s := NewSet[int]()
	if !s.Add(1) {
		t.Errorf("Unable to add int element to set")
	}
}

func TestSetAddString(t *testing.T) {
	s := NewSet[string]()
	if !s.Add("1") {
		t.Errorf("Unable to add int element to set")
	}
}

func TestSetContainsInt(t *testing.T) {
	s := NewSet[int]()
	got := s.Contains(0)
	want := false
	if got != want {
		t.Errorf("Got len %t, want len %t\n", got, want)
	}
}

func TestSetLengthInt(t *testing.T) {
	s := NewSet[int]()
	got := s.Len()
	want := 0

	if got != want {
		t.Errorf("Got len %d, want len %d\n", got, want)
	}
}

func TestSetEraseInt(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Erase(1)
	if s.Len() != 0 {
		t.Errorf("Unable to erase element from set")
	}
	s.Add(1)
	if s.Erase(2) {
		t.Errorf("Erasing element not present in set")
	}
}
