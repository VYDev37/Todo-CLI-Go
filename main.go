package main

import (
	"flag"
	"fmt"
	"os"
	model "todo-cli-go/todo"
)

func main() {
	manager := model.TaskManager{}
	manager.SetFile("tasks.json")

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [command name] [flags...]")
		return
	}

	err := manager.Load()
	if err != nil {
		fmt.Printf("Error at loading: %v.", err)
		return
	}
	taskNamePtr := flag.String("name", "", "Name")
	taskDuePtr := flag.String("due", "DD/MM/YY", "Due date")
	taskIdPtr := flag.Int("id", 0, "Task ID")

	flag.CommandLine.Parse(os.Args[2:]) // skip ke args > 2

	commandName := os.Args[1]
	taskName := *taskNamePtr
	taskDue := *taskDuePtr
	taskId := int16(*taskIdPtr)

	CommandHandler(commandName, &manager, CommandArgs{
		Name: taskName,
		Due:  taskDue,
		ID:   taskId,
	})
}
