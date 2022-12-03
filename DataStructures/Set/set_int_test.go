package Set

import (
	"testing"
)

func TestSetCreation(t *testing.T) {
	s := NewIntSet()

	if s == nil {
		t.Errorf("Got an empty set")
	}
}

func TestSetLength0(t *testing.T) {
	s := NewIntSet()
	got := s.Len()
	want := 0

	if got != want {
		t.Errorf("Got len %d, want len %d\n", got, want)
	}
}

func TestContainsLength0(t *testing.T) {
	s := NewIntSet()
	got := s.Contains(0)
	want := false
	if got != want {
		t.Errorf("Got len %t, want len %t\n", got, want)
	}
}

func TestAdd0(t *testing.T) {
	s := NewIntSet()
	got := s.Add(0)
	want := true
	if got != want {
		t.Errorf("Got len %t, want len %t\n", got, want)
	}
}

func TestSet(t *testing.T) {
	s := NewIntSet()
	got := s.Len()
	want := 0
	if got != want {
		t.Errorf("Got len %d, want len %d\n", got, want)
	}
	val := 0
	s.Add(val)
	got = s.Len()
	want = 1
	if got != want {
		t.Errorf("Got len %d, want len %d\n", got, want)
	}
	if s.Contains(0) == false {
		t.Errorf("Inserted value %d is not present in the set", val)
	}
	s.Add(1)
	s.Add(-1)
	if s.Len() != 3 {
		t.Errorf("Inserted Elements size does not match len")
	}

	if !s.Contains(1) || !s.Contains(-1) {
		t.Errorf("Inserted elements not present in the set")
	}
}

func TestDelete(t *testing.T) {
	s := NewIntSet()
	s.Add(1)
	s.Erase(1)
	if s.Len() != 0 {
		t.Errorf("Unable to delete the element")
	}
}
