package todolist

import (
	"errors"

	"github.com/mohamed-Dhia/to-go/counter"
	"github.com/mohamed-Dhia/to-go/todo"
)

type TodoList struct {
	counter *counter.Counter
	todos   []*todo.Todo
}

func New() *TodoList {
	return &TodoList{counter.New(), []*todo.Todo{}}
}

func (l *TodoList) Add(task string) *todo.Todo {
	newTodo := todo.New(l.counter.GetNextUniqueId(), task)
	l.todos = append(l.todos, newTodo)
	return newTodo
}

func (l *TodoList) String() (res string) {
	for _, todo := range l.todos {
		res += todo.String() + "\n"
	}
	return
}

func (l TodoList) FindTodoById(id int) (*todo.Todo, error) {
	for _, t := range l.todos {
		if t.Id == id {
			return t, nil
		}
	}
	return nil, errors.New("not found")
}

func (l *TodoList) ClearCompletedTodos() []*todo.Todo {
	var newTodos []*todo.Todo
	for _, todo := range l.todos {
		if !todo.Complete {
			newTodos = append(newTodos, todo)
		}
	}
	l.todos = newTodos
	return newTodos
}
