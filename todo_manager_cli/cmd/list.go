package cmd

import (
	"fmt"
	"os"

	"github.com/ILLIDOM/gophercises/todo_manager_cli/storage"
	"github.com/spf13/cobra"
)

var all bool //flag to show also completed todos

func init() {
	listCmd.Flags().BoolVar(&all, "all", false, "Flag to show also completed todos")
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print all todos",
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := storage.GetTodos()
		if err != nil {
			fmt.Println("ERROR: could no fetch todos:", err)
			os.Exit(1)
		}
		storage.PrintTodos(todos, all)
	},
}
