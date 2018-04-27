package queue

import (
	"testing"
)

func TestLen(t *testing.T) {
	q := New(10)
	if !q.isEmpty() {
		t.FailNow()
	}

	q.Push(1)
	if q.Len() != 1 {
		t.FailNow()
	}
}

func TestPush(t *testing.T) {
	q := New(10)
	q.Push(1)
	if value, err := q.Pop(); value != 1 && err != nil {
		t.FailNow()
	}
}

func TestPop(t *testing.T) {
	q := New(10)
	q.Push(1)
	if value, err := q.Pop(); value != 1 && err != nil {
		t.FailNow()
	}

}

func BenchmarkPush(b *testing.B) {
	q := New(b.N)
	b.ResetTimer() // 这里开始计时，或者使用 StopTimer/StartTimer

	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	q := New(b.N)

	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
	b.ResetTimer() // 这里开始计时，或者使用 StopTimer/StartTimer

	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}
