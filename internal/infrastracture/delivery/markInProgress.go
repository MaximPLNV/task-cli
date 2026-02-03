/*
Copyright © 2026 Maksim Palynov <m.palynov@gmail.com>
*/
package cmd

import (
	"fmt"
	"strconv"
	"task-cli/internal/interfaces"

	"github.com/spf13/cobra"
)

func NewMarkInProgress(uc interfaces.UseCase) *cobra.Command {
	return &cobra.Command{
		Use:   "mark-in-progress <id>",
		Short: "Update task's status to \"in-progress\"",
		Long: `
		Use this command to update task's status and save it to JSON file
		Syntax: task-cli mark-in-progress <task id>
		`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			idstr := args[0]
			id, convErr := strconv.Atoi(idstr)

			if convErr != nil {
				fmt.Println("Incorrect \"id\" format. Should be integer")
				return
			}

			updErr := uc.UpdateStatus(id, "in-progress")

			if updErr != nil {
				fmt.Println(updErr)
				return
			}
			fmt.Println("Status has been update to \"in-progress\"")
		},
	}
}
