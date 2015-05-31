package collections

type Set struct {
	data   map[interface{}]bool
	Values []interface{}
}

// add a value into set,
// if this set already has the value it will return false,
// else return true
func (s *Set) Add(val interface{}) bool {
	if s.data[val] {
		return false
	}

	s.data[val] = true
	s.Values = append(s.Values, val)
	return true
}
