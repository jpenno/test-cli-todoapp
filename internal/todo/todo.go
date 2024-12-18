package todo

import "fmt"

type Todo struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func New(id int, task string, done bool) Todo {
	return Todo{
		Id:   id,
		Task: task,
		Done: done,
	}
}

func PrintTodos(todos []Todo) {
	for _, t := range todos {
		t.Print()
	}
}

func (t Todo) Print() {
	fmt.Printf("ID: %v\n", t.Id)
	fmt.Printf("task: %v\n", t.Task)
	if t.Done {
		fmt.Println("Done")
	} else {
		fmt.Println("Not done")
	}
}
