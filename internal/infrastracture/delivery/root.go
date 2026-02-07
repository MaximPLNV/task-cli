/*
Copyright © 2026 Maksim Palynov <m.palynov@gmail.com>
*/
package cmd

import (
	"os"
	"task-cli/internal/interfaces"

	"github.com/spf13/cobra"
)

func NewRootCmd(uc interfaces.UseCase) *cobra.Command {

	rootCmd := &cobra.Command{
		Use:   "task-cli",
		Short: "A simple application to track your daily tasks",
		Long: `
		A task-cli is a pet project.
		Main purpose is tracking daily tasks
		Tasks records are stored in JSON file and can be created/updated/deleted
		`,
	}

	initSubcommands(rootCmd, uc)

	return rootCmd
}

func initSubcommands(cmd *cobra.Command, uc interfaces.UseCase) {
	for _, subcmd := range getSubCommands() {
		cmd.AddCommand(subcmd(uc))
	}
}

func getSubCommands() *[6]func(interfaces.UseCase) *cobra.Command {
	subcmds := [6](func(uc interfaces.UseCase) *cobra.Command){
		NewAddCmd,
		NewDeleteCmd,
		NewListCmd,
		NewMarkDoneCmd,
		NewMarkInProgress,
		NewUpdateCmd,
	}

	return &subcmds
}

func Execute(uc interfaces.UseCase) {
	root := NewRootCmd(uc)
	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}
