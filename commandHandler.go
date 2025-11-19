package main

import (
	"fmt"
	"strings"
	model "todo-cli-go/todo"
)

type CommandArgs struct {
	Name string
	Due  string
	ID   int16
}

func CommandHandler(commandName string, manager *model.TaskManager, args CommandArgs) {
	switch strings.ToLower(commandName) {
	case "--add", "add":
		err := manager.Add(args.Name, args.Due)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		fmt.Println("Task has been added to list.")
		manager.Print()
	case "--list", "list":
		manager.Print()
	case "--remove", "remove", "--delete", "delete":
		err := manager.Remove(args.ID)
		if err != nil {
			fmt.Printf("Error when trying to remove task #%d: %v.", args.ID, err)
			return
		}

		manager.Print()
	case "--update", "update", "--done", "done":
		err := manager.MarkasDone(args.ID)
		if err != nil {
			fmt.Printf("Error when trying to mark task #%d as done: %v.", args.ID, err)
			return
		}

		manager.Print()
	default:
		fmt.Println("Invalid command.")
	}
}
