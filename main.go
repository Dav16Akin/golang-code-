package main

import (
	"fmt"
)

func avgSlice64(slice []float64) (avg float64) {
	sum := 0.0
	switch len(slice) {
	case 0:
		avg = 0
	default:
		for _, v := range slice {
			sum += v
		}

		avg = sum / float64(len(slice))
	}
	return
}

func rightNumAsc(num1 , num2 int) (int, int) {
	if (num1 > num2) {
		return num2, num1
	}

	return num1, num2
}

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	if s.i >= len(s.data) {
		fmt.Println("Stack overflow")
		return
	}

	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() (int, bool) {
	if s.i == 0 {
		fmt.Println("Stack underflow")
		return 0 , false
	}
	s.i--
	return s.data[s.i], true
}

func printArgs(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func fibonacci(num int) []int {
	array := make([]int, num)

	array[0], array[1] = 1,1

	for i := 2; i < num; i++ {
		array[i] = array[i-1] + array[i-2]
	}

	return array
}

func Map(f func(int) int,l []int) []int {
	list := make([]int, len(l))

	for k,v := range l {
		list[k] = f(v)
	}

	return  list
}

func max(s []int) int {
	max := s[0]
	for _,v := range s {
		if v > max {
			max = v
		}
	}

	return max
}

func min(s []int) int {
	min := s[0]
	for _,v := range s {
		if v < min {
			min = v
		}
	}

	return min
}

func bubbleSort(l []int) {
	for i := 0; i < len(l)-1; i++ {
		for j := i + 1; j < len(l); j++ {
			if l[i] > l[j] {
				l[i], l[j] = l[j], l[i]
			}
		}
	}
}

func main() {
}
