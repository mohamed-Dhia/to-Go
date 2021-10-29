package todo

import (
	"fmt"
)

type Todo struct {
	Id       int
	Task     string
	Complete bool
}

func New(id int, task string) *Todo {
	return &Todo{Id: id, Task: task, Complete: false}
}

func (t *Todo) String() string {
	var status string

	if t.Complete {
		status = "completed"
	}
	status = "not completed yet!"
	return fmt.Sprintf("%d. %s is %s", t.Id, t.Task, status)
	// return strconv.Itoa(t.id) + ". " + t.task + " is " + status
}

func (t *Todo) ToggleComplete() {
	t.Complete = !t.Complete
}
