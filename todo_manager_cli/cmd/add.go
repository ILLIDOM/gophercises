package cmd

import (
	"errors"

	"github.com/ILLIDOM/gophercises/todo_manager_cli/storage"
	"github.com/spf13/cobra"
)

var (
	todoName string
)

func init() {
	RootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a todo",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires exactly one todo")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		todoName = args[0]
		storage.AddTodo(todoName)
	},
}
