package stack

import (
	"testing"
)

func TestLen(t *testing.T) {
	s := New(10)
	if !s.isEmpty() {
		t.FailNow()
	}

	s.Push(1)
	if s.Len() != 1 {
		t.FailNow()
	}
}

func TestPop(t *testing.T) {
	s := New(10)
	s.Push(1)
	if value, err := s.Pop(); err != nil && value != 1 {
		t.FailNow()
	}
}

func TestPush(t *testing.T) {
	s := New(10)
	s.Push(1)
	s.Push(2)
	if value, err := s.Pop(); err != nil && value != 2 {
		t.FailNow()
	}

}

func BenchmarkPush(b *testing.B) {
	s := New(10)
	b.ResetTimer() // 这里开始计时，或者使用 StopTimer/StartTimer

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	s := New(b.N)

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	b.ResetTimer() // 这里开始计时，或者使用 StopTimer/StartTimer

	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
