package main

import (
	"errors"
	"fmt"
	"os"
)

type quenu struct {
	maxSize int
	array   [4]int
	front   int
	real    int
}

func (a *quenu) AddQueue(val int) (err error) {
	if a.IsFull() {
		return errors.New("full")
	}
	a.array[a.real] = val
	a.real = (a.real + 1) % a.maxSize
	return
}

func (a *quenu) GetQueue() (val int, err error) {
	if a.IsEmpty() {
		return -1, errors.New("empty")
	}
	val = a.array[a.front]
	a.front = (a.front + 1) % a.maxSize
	return
}
func (a *quenu) ShowQueue() {
	if a.IsEmpty() {
		fmt.Println("empty")
		return
	}
	for i := a.front; i != a.real; i = (i + 1) % a.maxSize {
		fmt.Printf("array[%d]=%d\n", i, a.array[i])
	}
}

func (a *quenu) IsFull() bool {
	return (a.real+1)%a.maxSize == a.front
}

func (a *quenu) IsEmpty() bool {
	return a.front == a.real
}

func main() {
	qu := &quenu{
		maxSize: 4,
		front:   0,
		real:    0,
	}
	for {
		fmt.Printf("现在有 get add show exit 操作\n")
		fmt.Printf("请输入你要的操作:\n")
		var msg string
		var val int
		fmt.Scanf("%s", &msg)
		switch msg {
		case "add":
			fmt.Printf("请输入你要添加的元素:\n")
			fmt.Scanf("%d", &val)
			err := qu.AddQueue(val)
			if err != nil {
				fmt.Println(err)
			}
			qu.ShowQueue()
			break
		case "get":
			data, err := qu.GetQueue()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("出队的元素是%d\n", data)
			fmt.Printf("现在栈中元素是:\n")
			qu.ShowQueue()
			break
		case "show":
			qu.ShowQueue()
			break
		case "exit":
			os.Exit(0)

		}
	}
}
