package add

import (
	fileio "todoApp/internal/fileIO"
	"todoApp/internal/todo"
)

func AddTodo(task string) {
	todos := fileio.GetTodos("./assets/data.json")

	newTodo := todo.New(len(todos)+1, task, false)
	todos = append(todos, newTodo)

	fileio.SaveTodos("./assets/data.json", todos)
}
