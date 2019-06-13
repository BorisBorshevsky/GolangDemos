package main

import "fmt"

type Stack interface {
	Len() int
	Push(int)
	Pop() (int, error)
}

var EmptySliceErr error = fmt.Errorf("this is empty")

type SliceStack []int

func (s *SliceStack) Len() int {
	return len(*s)
}

func (s *SliceStack) Push(val int) {
	*s = append(*s, val)
}

func (s *SliceStack) Pop() (int, error) {
	if s.Len() == 0 {
		return 0, EmptySliceErr
	}
	result := (*s)[s.Len()-1]

	*s = (*s)[:s.Len()-1]

	return result, nil

}

func NewSliceStack() *SliceStack {
	return &SliceStack{}
}

func main() {
	var s Stack = NewSliceStack()
	fmt.Println(s)
	s.Push(1)
	s.Push(2)
	fmt.Println(s)
	if v, err := s.Pop(); err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("pop", v)
	}
	fmt.Println(s)

}
