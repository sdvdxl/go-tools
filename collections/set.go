package collections
import "fmt"

type Set struct {
	data   map[interface{}]bool
	values []interface{}
}

func (s Set) Size() int {
	return len(s.values)
}

func NewSet(size int, values ...interface{}) *Set {
	set := &Set{}
	set.data = make(map[interface{}]bool)
	for _, val := range values {
		set.Add(val)
	}

	return set
}

// add a value into set,
// if this set already has the value it will return false,
// else return true
func (s *Set) Add(val interface{}) bool {
	if s.data[val] {
		return false
	}

	s.data[val] = true
	s.values = append(s.values, val)
	return true
}

//return true if this set contains the value,
//otherwise will return false
func (s *Set) Contains(val interface{}) bool {
	if s.data[val] {
		return true
	}

	return false
}

// if contains the value then delete the value and return ture
// otherwise will return false
func (s *Set) Remove(val interface{}) bool {
	if s.data[val] {
		delete(s.data, val)
		tmpSet := make([]interface{}, 0, len(s.data))
		for k, _ := range s.data {
			tmpSet = append(tmpSet, k)
		}

		s.values = tmpSet
		return true
	}

	return false
}

// get values from the set
func (s *Set) Values() []interface{} {
	return s.values
}

//values len is 0 will return true
// otherwise will return false
func (s *Set) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Set) String() string {
	var result string
	for v,_ :=range s.data {
		result += fmt.Sprintf("%v", v)
	}

	return result
}