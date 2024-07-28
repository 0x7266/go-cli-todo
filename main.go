package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

func main() {
	// Todos := make([]Todo, 0)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What do you want to do next?\nlist\ncreate\nupdate\ndelete")

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		switch line[0] {
		case "list":
			break
		case "create":
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
