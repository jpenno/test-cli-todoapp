/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"todoApp/internal/add"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new todo",
	Long:  `add new todo`,
	Run:   addTodoCmd,
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("task", "t", "", "todo task")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addTodoCmd(cmd *cobra.Command, args []string) {
	task, err := cmd.Flags().GetString("task")
	if err != nil {
		panic(err)
	}

	add.AddTodo(task)
}
