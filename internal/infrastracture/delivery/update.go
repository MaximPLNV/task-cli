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

func NewUpdateCmd(uc interfaces.UseCase) *cobra.Command {
	return &cobra.Command{
		Use:   "update <id> <description>",
		Short: "Update existing task",
		Long: `
		Use this command to udapte task's description and save it to JSON file
		Syntax: task-cli update <task id> <string with new task description>
		`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			idstr, d := args[0], args[1]
			id, convErr := strconv.Atoi(idstr)

			if convErr != nil {
				fmt.Println("Incorrect \"id\" format. Should be integer")
				return
			}

			updErr := uc.Update(id, d)

			if updErr != nil {
				fmt.Println(updErr)
				return
			}
			fmt.Println("Task has been update")
		},
	}
}
