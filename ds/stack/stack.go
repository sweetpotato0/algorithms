package stack

import (
	"errors"
	"sync"
)

type Stack struct {
	sync.RWMutex
	data   []interface{}
	length int
	space  int
}

func New(space int) *Stack {
	if space <= 0 {
		panic(errors.New("space must be more than 0."))
	}
	return &Stack{
		data:   make([]interface{}, 0),
		length: 0,
		space:  space,
	}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) isEmpty() bool {
	return s.length == 0
}

func (s *Stack) Push(value interface{}) error {
	if s.length >= s.space {
		return errors.New("Stack is full.")
	}

	s.Lock()
	s.data = append(s.data, value)
	s.length++
	s.Unlock()

	return nil
}

func (s *Stack) Pop() (value interface{}, err error) {
	if s.isEmpty() {
		return nil, errors.New("Stack is empty.")
	}
	n := s.Len()

	s.Lock()
	value = s.data[n-1]
	s.data = s.data[:n-1]
	s.length--
	s.Unlock()

	return value, nil
}
