package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
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
	fmt.Println("What do you want to do next?\nlist\ncreate\nupdate\ndelete\n\n")

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		switch line[0] {
		case "list":
			fmt.Println("\n")
			for i := 0; i < len(Todos); i++ {
				fmt.Printf("%v\nID: %v\nStatus: %v\n\n", Todos[i].Description, Todos[i].Id, checkIfDone(Todos[i].IsDone))
			}
			break
		case "create":
			Todos = append(Todos, Todo{
				Id:          len(Todos) + 1,
				Description: strings.Join(line[1:], " "),
				IsDone:      false,
			})
			break
		case "update":
			break
		case "delete":
			break
		case "done":
			break
		}
	}
}
