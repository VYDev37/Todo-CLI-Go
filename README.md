# TODO-CLI-GO

## Description: Only a simple Todo CLI App
## Project only for my Go learning purpose.

## Commands
- `add` (aliases: `--add`) -> To add task into list.
  Params: name `(-name)`, due date `(-due)`
- `list` (aliases: `--list`) -> To display the list of tasks.
  Params: None.
- `remove` (aliases: `--remove`, `--delete`, `delete`) -> To remove the task from list.
  Params: task id `(-id)`
- `update` (aliases: `--update`, `--done`, `done`) -> To mark the task as done.
  Params: task id `(-id)`

## How to run:
1. Open Terminal
2. Type `go run . [command_name] -[param1] [value] -[param2] [value] ...`

## Examples of how to run:
1. add: `go run . --add -name "Finish Computer Science Assignment" -due "19/12/2025"`
2. list: `go run . --list`
3. remove: `go run . --remove -id 1`
4. update: `go run . --update -id 1`

### Next plan: Make backend version of this (REST + gRPC API), and make some needed changes and refactor.