package redis_lock

import (
	"fmt"
	"sync"
	"testing"
)

type inter interface{}
type Set struct {
	m map[inter]bool
	sync.RWMutex
}

func New() *Set {
	return &Set{
		m: map[inter]bool{},
	}
}
func (s *Set) Add(item inter) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func TestT2(t *testing.T) {
	set := New()
	set.Add("1")
	set.Add("2")
	fmt.Println("set:", set)
}
