package stack

import "fmt"

type Stack struct {
	i    int
	data [10]int
}

func (s *Stack) Push(k int) {
	if s.i >= len(s.data) {
		fmt.Println("Stack overflow")
		return
	}

	s.data[s.i] = k
	s.i++
}

func (s *Stack) Pop() (int, bool) {
	if s.i == 0 {
		fmt.Println("Stack underflow")
		return 0 , false
	}
	s.i--
	return s.data[s.i], true
}
