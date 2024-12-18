/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"todoApp/internal/get"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list of all todo's",
	Long:  `list tasks`,
	Run:   listTodoCmd,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntP("id", "i", 0, "get task by id")
	listCmd.Flags().BoolP("all", "a", false, "list all tasks")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listTodoCmd(cmd *cobra.Command, args []string) {
	id, err := cmd.Flags().GetInt("id")
	if err != nil {
		panic(err)
	}

	listAll, err := cmd.Flags().GetBool("all")
	if err != nil {
		panic(err)
	}

	if listAll {
		todos := get.GetTodos()
		for _, t := range todos {
			t.Print()
		}
	}

	if id != 0 {
		todo := get.GetTodoById(id)
		todo.Print()
	}
}
