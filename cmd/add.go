package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/connor-cahill/goDoList/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong: ", err.Error())
			os.Exit(1)
			return
		}
		fmt.Printf("Added \"%s\" to your task list!\n", task)

	},
}

//* init function runs before anything else in the file
func init() {
	RootCmd.AddCommand(addCmd) //	adds the add command as a child command of Root command
}
