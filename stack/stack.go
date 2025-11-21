package main

import (
	"fmt"
	"os"
)

type Stack struct {
	maxSize int
	array   []int
	top     int
}

func (s *Stack) Push(val int) {
	if s.IsFull() {
		panic("栈已满")
	}
	s.top++
	s.array[s.top] = val
}
func (s *Stack) Pop() (val int) {
	if s.IsEmpty() {
		panic("栈已空")
	}
	val = s.array[s.top]
	s.top--
	return
}
func (s *Stack) IsFull() bool {
	return s.top == s.maxSize-1
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) List() {
	if s.IsEmpty() {
		panic("栈已空")
	}
	for i := s.top; i >= 0; i-- {
		fmt.Println(s.array[i])
	}
}

func main() {
	a := &Stack{
		maxSize: 4,
		array:   make([]int, 4),
		top:     -1,
	}
	for {
		fmt.Printf("现在有 pop push list exit 操作\n")
		fmt.Printf("请输入你要的操作:\n")
		var msg string
		var val int
		fmt.Scanf("%s", &msg)
		switch msg {
		case "push":
			fmt.Printf("请输入你要添加的元素:\n")
			fmt.Scanf("%d", &val)
			a.Push(val)
			a.List()
			break
		case "pop":
			s := a.Pop()
			fmt.Printf("出栈的元素是%d\n", s)
			fmt.Printf("现在栈中元素是:\n")
			a.List()
			break
		case "list":
			a.List()
			break
		case "exit":
			os.Exit(0)

		}
	}
}
