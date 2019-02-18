package cmd

import "github.com/spf13/cobra"

//RootCmd is the root command for our CLI app
var RootCmd = &cobra.Command{
	Use:   "godo",
	Short: "godo is a CLI task manager",
}
