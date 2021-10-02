package main

import "fmt"

func main() {
	myTodoList := &TodoList{
		counter: Counter{},
	}
	myTodoList.add("brush teeth")
	myTodoList.add("do stuff")
	myTodoList.add("do them more")
	fmt.Println(myTodoList.display())
}
