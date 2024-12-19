package delete

import (
	"fmt"
	"os"
	fileio "todoApp/internal/fileIO"
	"todoApp/internal/get"
	"todoApp/internal/todo"
)

func remove(slice []todo.Todo, s int) []todo.Todo {
	return append(slice[:s], slice[s+1:]...)
}

func DeleteById(id int) {
	todos := get.GetTodos()

	fmt.Println("Deleting")
	todos[id-1].Print(os.Stdout)

	todos = remove(todos, id-1)
	for i := range todos {
		todos[i].Id = i + 1
	}

	fileio.SaveTodos("./assets/data.json", todos)
}
