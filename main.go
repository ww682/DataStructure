package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array   [4]int
	front   int //头部
	real    int //尾部
}

func (a *Queue) AddQueue(val int) (err error) {
	if a.real == a.maxSize-1 {
		return errors.New("full")
	}
	a.real++
	a.array[a.real] = val
	return
}

func (a *Queue) ShowQueue() {
	for i := a.front + 1; i <= a.real; i++ {
		fmt.Printf("array[%d]=%d\n", i, a.array[i])
	}
}

func main() {
	var key string
	var val int

	queue := &Queue{
		maxSize: 5,
		front:   -1,
		real:    -1,
	}
	for {
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("success")
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
