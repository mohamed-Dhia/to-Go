package main

import (
	"fmt"

	todolist "github.com/mohamed-Dhia/to-go/todoList"
)

func main() {
	myTodoList := todolist.New()
	myTodoList.Add("brush teeth")
	myTodoList.Add("do stuff")
	myTodoList.Add("do them more")
	fmt.Printf("%s\n", myTodoList)
}
