package cmd

import (
	"fmt"
	"strconv"
	"task-cli/internal/domain/interfaces"

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
		Run:  getDeleteRunFunc(uc),
	}
}

func getDeleteRunFunc(uc interfaces.UseCase) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		idstr := args[0]
		id, convErr := strconv.Atoi(idstr)

		if convErr != nil {
			fmt.Println("Incorrect \"id\" format. Should be integer")
			return
		}

		if delErr := uc.Delete(id); delErr != nil {
			fmt.Println(delErr)
			return
		}

		fmt.Println("Task has been deleted")
	}
}
