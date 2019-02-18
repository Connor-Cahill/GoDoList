package cmd

import (
	"fmt"
	"strconv"

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
		fmt.Println(ids) //	for now, just print ids
	},
}

//* init function runs before anything else in the file
func init() {
	RootCmd.AddCommand(doCmd) //	adds the add command as a child command of Root command
}
