package get

import (
	"fmt"
	"todoApp/internal/fileIO"
)

func GetTodos() {
	todos := fileio.GetTodos("./assets/data.json")

	for _, t := range todos {
		t.Print()
		fmt.Println()
	}
}
