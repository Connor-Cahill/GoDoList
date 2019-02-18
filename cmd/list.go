package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists out all tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		//	stuff
	},
}

//* init function runs before anything else in the file
func init() {
	RootCmd.AddCommand(listCmd) //	adds the add command as a child command of Root command
}
