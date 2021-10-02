package main

import "strconv"

type Todo struct {
	id       int
	task     string
	complete bool
}

func createTodo(id int, task string) Todo {
	return Todo{id: id, task: task, complete: false}
}

func (t Todo) Display() string {
	var status string

	if t.complete {
		status = "completed"
	}
	status = "not completed yet !"
	return strconv.Itoa(t.id) + ". " + t.task + " is " + status
}

func (t *Todo) ToggleComplete() {
	t.complete = !t.complete
}
