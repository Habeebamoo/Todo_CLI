package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
)

type Todo struct {
	Name string `json:"name"`
	Completed bool `json:"completed"`
}

func loadTodo() ([]Todo, error) {
	todosJson, err := os.ReadFile("todos.json")
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create("todos.json")
			check(err)
		}
		return nil, errors.New("error loading todos.json file")
	}

	var todos []Todo
	err = json.Unmarshal(todosJson, &todos)
	if err != nil {
		return nil, errors.New("error reading todos.json file")
	}
	return todos, nil
}

func saveTodo(todos []Todo) error {
	todosJson, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return errors.New("error formatting todos file")
	}

	err = os.WriteFile("todos.json", todosJson, 0644)
	if err != nil {
		return errors.New("error saving todos.json file")
	}
	return nil
}

func addTodo(todo string) error {
	if todo == "" {
		return errors.New("no todo given")
	}

	todos, err := loadTodo()
	check(err)

	for i, _ := range todos {
		if todos[i].Name == todo {		
			return errors.New("todo already exists")
		}
	}

	todos = append(todos, Todo{todo, false})
	err = saveTodo(todos)
	check(err)

	fmt.Println("todo added successfully")
	return nil
}

func completeTodo(todo string) error {
	if todo == "" {
		return errors.New("no todo given")
	}

	todos, err := loadTodo()
	check(err)

	for i, _ := range todos {
		if todos[i].Name == todo {		
			todos[i].Completed = true

			err = saveTodo(todos)
			check(err)
	
			fmt.Println("todo has been marked")
			return nil
		}
	}

	return errors.New("todo was not found")
}

func removeTodo(todo string) error {
	if todo == "" {
		return errors.New("no todo given")
	}

	todos, err := loadTodo()
	check(err)

	for i, _ := range todos {
		if todos[i].Name == todo {
			todos = append(todos[:i], todos[i+1:]... )

			err := saveTodo(todos)
			check(err)

			fmt.Println("todo has been deleted")
			return nil
		}
	}

	return errors.New("todo was not found")
}

func listTodo() error {
	todos, err := loadTodo()
	check(err)

	for _, todo := range todos {
		fmt.Printf("\n********************\nName: %s\nCompleted: %t\n", todo.Name, todo.Completed)
	}
	return nil
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	var t string
	var todo string

	flag.StringVar(&t, "t", "", "add, complete, delete, list")
	flag.StringVar(&todo, "v", "", "Enter todo")

	flag.Parse()

	switch t {
	case "add":
		err := addTodo(todo)
		check(err)
	case "complete":
		err := completeTodo(todo)
		check(err)
	case "delete":
		err := removeTodo(todo)
		check(err)
	case "list":
		err := listTodo()
		check(err)
	default:
		fmt.Println("Invalid Command")
	}
}