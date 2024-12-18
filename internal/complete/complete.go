package complete

import (
	"fmt"
	fileio "todoApp/internal/fileIO"
	"todoApp/internal/get"
)

func CompleteById(id int) {
	todos := get.GetTodos()

	todos[id-1].Done = true
	fmt.Println("Task done")
	fmt.Println(todos[id-1].Task)

	fileio.SaveTodos("./assets/data.json", todos)
}
