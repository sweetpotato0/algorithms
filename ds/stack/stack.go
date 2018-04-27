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
		data:   make([]interface{}, space),
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
	defer s.Unlock()
	s.Lock()

	if s.length >= s.space {
		return errors.New("Stack is full.")
	}
	s.data[s.length] = value
	s.length++

	return nil
}

func (s *Stack) Pop() (value interface{}, err error) {
	defer s.Unlock()
	s.Lock()

	if s.isEmpty() {
		return nil, errors.New("Stack is empty.")
	}
	value = s.data[s.length-1]
	s.data = s.data[:s.length-1]
	s.length--

	return value, nil
}
