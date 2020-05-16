package utils

type unit struct{}
type StringSet map[string]unit

func (s StringSet) Add(x string) {
	s[x] = unit{}
}

func (s StringSet) Remove(x string) {
	delete(s, x)
}

func (s StringSet) Contains(x string) bool {
	_, ok := s[x]
	return ok
}
