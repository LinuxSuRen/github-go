package cmd

import "github.com/spf13/cobra"

// NewRoot returns the root command of github-go
func NewRoot() *cobra.Command {
	return &cobra.Command{
		Use: "github-go",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("hello")
		},
	}
}
