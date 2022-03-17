package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	top int
	arr []byte
}

func (s *Stack) push(e byte) {
	if len(s.arr) == s.top {
		s.arr = append(s.arr, e)
		s.top++
		return
	}
	s.arr[s.top] = e
	s.top++
}

func (s *Stack) pop() (byte, error) {
	if s.top == 0 {
		return 0, errors.New("Stack Empty")
	}
	s.top--
	v := s.arr[s.top]
	return v, nil
}

func match(str string) bool {
	stk := Stack{}
	for _, v := range str {
		if v == '(' {
			stk.push(byte(v))
		} else {
			if _, err := stk.pop(); err != nil {
				return false
			}
		}
	}
	if stk.top > 0 {
		return false
	}
	return true
}

func main() {
	input := ")()()("
	if match(input) {
		fmt.Println("matched")
	} else {
		fmt.Println("not match")
	}
}
