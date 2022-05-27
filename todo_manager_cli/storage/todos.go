package storage

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

var todoBucket = []byte("todos") // in memory bucket holding todos after fetch from db
var db *bolt.DB                  // db connection

type Todo struct {
	Id        int    // auto generated id used as key in database
	Name      string // name of todo used as value in database
	Completed bool   // flag showing if a todo is completed or not
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(todoBucket)
		return err
	})
}

func GetTodos() ([]Todo, error) {
	var todos []Todo

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(todoBucket)
		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			todo := Todo{}
			err := json.Unmarshal(v, &todo)
			if err != nil {
				fmt.Println("Error decoding todo", err)
			}
			todos = append(todos, todo)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return todos, nil
}

// prints todos, all can toggle to show also completed todos
func PrintTodos(todos []Todo, all bool) {
	if len(todos) == 0 {
		fmt.Println("You don't have any todos!")
		return
	}
	fmt.Println("Open todos:")
	for index, todo := range todos {
		var status string
		if todo.Completed {
			status = "completed"
		} else {
			status = "not completed"
		}
		if all || !todo.Completed {
			fmt.Printf("%v. %s: %v\n", index+1, todo.Name, status)
		}
	}
}

func AddTodo(todo string) (int, error) {
	var id int
	var new_todo = Todo{}

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(todoBucket)
		id64, _ := bucket.NextSequence()
		id = int(id64)
		key := itob(id) //convert id to byte array
		new_todo.Id = id
		new_todo.Name = todo
		encoded_todo, err := json.Marshal(new_todo)
		if err != nil {
			fmt.Println("Error encoding todo", err)
		}
		return bucket.Put(key, encoded_todo)
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

// marks a todo as done
func MarkTodoAsDone(todo Todo) error {
	todo.Completed = true

	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(todoBucket)
		encoded_todo, err := json.Marshal(todo)
		if err != nil {
			fmt.Println("Error encoding todo", err)
		}
		return bucket.Put(itob(todo.Id), encoded_todo)
	})
}

// removes a todo from the database
func RemoveTodo(todoId int) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(todoBucket)
		return bucket.Delete(itob(todoId))
	})
}

// converts integer to bytes array
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// converts byte array to integer
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
