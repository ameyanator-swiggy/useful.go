package Set

type setInt struct {
	m map[int]bool
}

func NewIntSet() *setInt {
	return &setInt{m: map[int]bool{}}
}

func (s *setInt) Contains(x int) bool {
	_, present := s.m[x]
	return present
}

func (s *setInt) Add(x int) bool {
	if s.Contains(x) {
		return false
	}
	s.m[x] = true
	return true
}

func (s *setInt) Len() int {
	return len(s.m)
}

func (s *setInt) Erase(x int) bool {
	if !s.Contains(x) {
		return false
	}
	delete(s.m, x)
	return true
}
