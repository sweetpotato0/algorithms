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
		data:   make([]interface{}, 0),
		length: 0,
		space:  space,
	}
}

func (q *Queue) Len() int {
	return q.length
}

func (q *Queue) isEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Push(value interface{}) error {
	if q.length >= q.space {
		return errors.New("Queue is full.")
	}

	q.Lock()
	q.data = append(q.data, value)
	q.length++
	q.Unlock()

	return nil
}

func (q *Queue) Pop() (value interface{}, err error) {

	if q.isEmpty() {
		return nil, errors.New("Queue is empty.")
	}

	q.Lock()
	value, q.data = q.data[0], q.data[:1]
	q.length--
	q.Unlock()

	return value, nil
}
