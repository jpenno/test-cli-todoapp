package get

import (
	fileio "todoApp/internal/fileIO"
	"todoApp/internal/todo"
)

func GetTodos() []todo.Todo {
	todos := fileio.GetTodos("./assets/data.json")
	return todos
}

func GetTodoById(id int) todo.Todo {
	todos := GetTodos()
	return todos[id-1]
}
