package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		//	stuff
	},
}

//* init function runs before anything else in the file
func init() {
	RootCmd.AddCommand(addCmd) //	adds the add command as a child command of Root command
}
