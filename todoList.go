package main

import "errors"

type TodoList struct {
	counter Counter
	todos   []Todo
}

func (l *TodoList) Add(task string) Todo {
	newTodo := createTodo(l.counter.GetNextUniqueId(), task)
	l.todos = append(l.todos, newTodo)
	return newTodo
}

func (l TodoList) Display() (res string) {
	for _, todo := range l.todos {
		res += todo.Display() + "\n"
	}
	return
}

func (l TodoList) FindTodoById(id int) (Todo, error) {
	for _, t := range l.todos {
		if t.id == id {
			return t, nil
		}
	}
	return Todo{}, errors.New("not found")
}

func (l *TodoList) ClearCompletedTodos() []Todo {
	var newTodos []Todo
	for _, todo := range l.todos {
		if !todo.complete {
			newTodos = append(newTodos,todo)
		}
	}
	l.todos = newTodos
	return newTodos
}