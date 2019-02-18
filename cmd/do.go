package cmd

import (
	"fmt"
	"strconv"

	"github.com/connor-cahill/goDoList/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg) //	converts the arg to an int
			if err != nil {
				fmt.Println("Failed to parse the argument: ", arg)
				break
			}
			ids = append(ids, id) //	append new id into our slice of ids
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number: ", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as complete. Error: %s\n", id, err)
			} else {
				fmt.Printf(" Marked \"%d\" as completed.\n", id)
			}

		}

	},
}

//* init function runs before anything else in the file
func init() {
	RootCmd.AddCommand(doCmd) //	adds the add command as a child command of Root command
}
