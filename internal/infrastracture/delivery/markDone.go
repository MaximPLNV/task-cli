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

func NewMarkDoneCmd(uc interfaces.UseCase) *cobra.Command {
	return &cobra.Command{
		Use:   "mark-done <id>",
		Short: "Update task's status to \"done\"",
		Long: `
		Use this command to update task's status and save it to JSON file
		Syntax: task-cli mark-done <task id>
		`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			idstr := args[0]
			id, convErr := strconv.Atoi(idstr)

			if convErr != nil {
				fmt.Println("Incorrect \"id\" format. Should be integer")
				return
			}

			updErr := uc.UpdateStatus(id, "done")

			if updErr != nil {
				fmt.Println(updErr)
				return
			}
			fmt.Println("Status has been update to \"done\"")
		},
	}
}
