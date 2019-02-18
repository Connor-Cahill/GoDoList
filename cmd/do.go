package cmd

import (
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete.",
	Run: func(cmd *cobra.Command, args []string) {
		//	stuff
	},
}

//* init function runs before anything else in the file
func init() {
	RootCmd.AddCommand(doCmd) //	adds the add command as a child command of Root command
}
