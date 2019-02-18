package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args) // prints out everything after command
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to your task list!\n", task)

	},
}

//* init function runs before anything else in the file
func init() {
	RootCmd.AddCommand(addCmd) //	adds the add command as a child command of Root command
}
