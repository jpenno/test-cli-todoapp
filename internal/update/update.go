package update

import (
	fileio "todoApp/internal/fileIO"
	"todoApp/internal/get"
)

func UpdateTodoById(id int, update string) {

	todos := get.GetTodos()

	// fmt.Println("Deleting")
	todos[id-1].Task = update

	fileio.SaveTodos("./assets/data.json", todos)
}
