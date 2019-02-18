package cmd

import (
	"fmt"
	"os"

	"github.com/connor-cahill/goDoList/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists out all tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong: ", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You currently have no tasks! 😤")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

//* init function runs before anything else in the file
func init() {
	RootCmd.AddCommand(listCmd) //	adds the add command as a child command of Root command
}
