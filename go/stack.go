package main

import "errors"

type Stack struct {
	data []float64
}

func (s *Stack) Push(i float64) {
	s.data = append(s.data, i)
}

func (s *Stack) Shift() (float64, error) {
	if len(s.data) <= 0 {
		return 0, errors.New("Stack is empty")
	}
	item := s.data[0]
	if len(s.data) > 1 {
		s.data = s.data[1:]
	} else {
		s.data = []float64{}
	}
	return item, nil
}
