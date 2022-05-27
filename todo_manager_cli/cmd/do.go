package cmd

import (
	"fmt"
	"strconv"

	"github.com/ILLIDOM/gophercises/todo_manager_cli/storage"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "mark todo as done",
	Run: func(cmd *cobra.Command, args []string) {
		var ids_to_remove []int

		for _, arg := range args {
			todoId, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed parsing argument", err)
			} else {
				ids_to_remove = append(ids_to_remove, todoId)
			}
		}

		todos, err := storage.GetTodos()
		if err != nil {
			fmt.Println("Couldn't fetch todos", err)
			return
		}

		for _, id := range ids_to_remove {
			if id <= 0 || id > len(todos) {
				fmt.Println("Invalid todo number:", id)
				continue
			}
			todo := todos[id-1]
			err := storage.RemoveTodo(todo.Id)
			if err != nil {
				fmt.Printf("Failed to mark %d as completed. Error %v\n", id, err)
			} else {
				fmt.Printf("Marked %d as completed. \n", id)
			}
		}
	},
}
