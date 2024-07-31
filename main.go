package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Todo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

func checkIfDone(status bool) string {
	if status {
		return "DONE"
	}
	return "NOT DONE"
}

func main() {
	Todos := make([]Todo, 0)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What do you want to do next?\nlist\ncreate\nupdate\ndelete\n> ")

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		switch line[0] {
		case "list":
			for _, todo := range Todos {
				fmt.Println("----------")
				statusLine := "[ ] "
				if todo.IsDone {
					statusLine = "[X] "
				}
				fmt.Printf("%v%v\nid: %v\n", statusLine, todo.Description, todo.Id)
			}
			break
		case "create":
			Todos = append(Todos, Todo{
				Id:          len(Todos) + 1,
				Description: strings.Join(line[1:], " "),
				IsDone:      false,
			})
			fmt.Println("Created TODO item")
			break
		case "update":
			break
		case "delete":
			id, err := strconv.Atoi(line[1])
			if err != nil {
				panic(err)
			}
			for i, todo := range Todos {
				if todo.Id == id {
					Todos = slices.Delete(Todos, i, i+1)
				}
			}
			break
		case "done":
			id, err := strconv.Atoi(line[1])
			if err != nil {
				panic(err)
			}
			for i, todo := range Todos {
				if todo.Id == id {
					Todos[i].IsDone = !Todos[i].IsDone
				}
			}
			break
		}
		fmt.Print("> ")
	}
}
