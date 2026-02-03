/*
Copyright © 2026 Maksim Palynov <m.palynov@gmail.com>
*/
package cmd

import (
	"fmt"
	"task-cli/internal/interfaces"

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
		Run: func(cmd *cobra.Command, args []string) {
			d := args[0]
			err := uc.Add(d)

			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Task has been created")
		},
	}
}
