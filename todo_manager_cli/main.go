package main

import (
	"path/filepath"

	"github.com/ILLIDOM/gophercises/todo_manager_cli/cmd"
	"github.com/ILLIDOM/gophercises/todo_manager_cli/storage"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "todos.db")
	storage.Init(dbPath)
	cmd.RootCmd.Execute()
}
