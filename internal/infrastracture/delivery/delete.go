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

func NewDeleteCmd(uc interfaces.UseCase) *cobra.Command {
	return &cobra.Command{
		Use:   "delete <id>",
		Short: "Delete existing task",
		Long: `
		Use this command to delete task record from JSON file
		Syntax: task-cli delete <task id>
		`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			idstr := args[0]
			id, convErr := strconv.Atoi(idstr)

			if convErr != nil {
				fmt.Println("Incorrect \"id\" format. Should be integer")
				return
			}

			delErr := uc.Delete(id)

			if delErr != nil {
				fmt.Println(delErr)
				return
			}
			fmt.Println("Task has been deleted")
		},
	}
}
