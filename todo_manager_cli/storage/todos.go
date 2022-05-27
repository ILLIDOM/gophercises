package storage

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

var todoBucket = []byte("todos") // in memory bucket holding todos after fetch from db
var db *bolt.DB                  // db connection

type Todo struct {
	Id   int    // auto generated id used as key in database
	Name string // name of todo used as value in database
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
			todos = append(todos, Todo{
				Id:   btoi(k),
				Name: string(v),
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func PrintTodos(todos []Todo) {
	if len(todos) == 0 {
		fmt.Println("You don't have any todos!")
		return
	}
	fmt.Println("Open todos:")
	for index, todo := range todos {
		fmt.Printf("%v. %s\n", index+1, todo.Name)
	}
}

func AddTodo(todo string) (int, error) {
	var id int

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(todoBucket)
		id64, _ := bucket.NextSequence()
		id = int(id64)
		key := itob(id)
		return bucket.Put(key, []byte(todo))
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

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
