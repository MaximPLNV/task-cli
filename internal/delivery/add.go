package cmd

import (
	"fmt"
	"task-cli/internal/domain/interfaces"

	"github.com/spf13/cobra"
)

func NewAddCmd(uc interfaces.UseCase) *cobra.Command {

	return &cobra.Command{
		Use:   "add <description>",
		Short: "Create new task",
		Long: `
		Use this command to create task record and save it to JSON file
		Syntax: task-cli add <string with task description>
		`,
		Args: cobra.ExactArgs(1),
		Run:  getAddRunFunc(uc),
	}
}

func getAddRunFunc(uc interfaces.UseCase) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		d := args[0]

		if err := uc.Add(d); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Task has been created")
	}
}
