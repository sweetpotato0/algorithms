package queue

import (
	"errors"
	"sync"
)

type Queue struct {
	sync.RWMutex
	data   []interface{}
	length int
	space  int
}

func New(space int) *Queue {
	if space <= 0 {
		panic(errors.New("space must be more than 0."))
	}
	return &Queue{
		data:   make([]interface{}, space),
		length: 0,
		space:  space,
	}
}

func (q *Queue) Len() int {
	return q.length
}

func (q *Queue) isEmpty() bool {
	return q.length == 0
}

func (q *Queue) Push(value interface{}) error {
	defer q.Unlock()
	q.Lock()

	if q.length >= q.space {
		return errors.New("Queue is full.")
	}
	q.data[q.length] = value
	q.length++

	return nil
}

func (q *Queue) Pop() (value interface{}, err error) {
	defer q.Unlock()
	q.Lock()

	if q.isEmpty() {
		return nil, errors.New("Queue is empty.")
	}
	value, q.data = q.data[0], q.data[:1]
	q.length--

	return value, nil
}
