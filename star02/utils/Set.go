package utils

import(
	"sync"
)

type Set struct{
	m map[interface{}]bool
	sync.RWMutex
}

func NewSet(m map[interface{}]bool) *Set{
	return &Set{
		m: m,
	}
}

func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}
  
func (s *Set) Remove(item int) {
	s.Lock()
	s.Unlock()
	delete(s.m, item)
}
  
func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}
  
func (s *Set) Len() int {
	return len(s.List())
}
  
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[interface{}]bool{}
}
  
func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
	  return true
	}
	return false
}
  
func (s *Set) List() []interface{} {
	s.RLock()
	defer s.RUnlock()
	list := make([]interface{}, 0)
	for item := range s.m {
	  list = append(list, item)
	}
	return list
}

func New() *Set {
	return &Set{
		m: make(map[interface{}]bool, 0),
	}
}

