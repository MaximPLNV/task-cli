package cmd

import (
	"fmt"
	"task-cli/internal/domain/interfaces"

	"github.com/spf13/cobra"
)

func NewListCmd(uc interfaces.UseCase) *cobra.Command {
	return &cobra.Command{
		Use:   "list [status]",
		Short: "Get list of tasks",
		Long: `
		Use this command to get list of tasks according to task's status
		Syntax: task-cli list [status]
		`,
		Args: cobra.MaximumNArgs(1),
		Run:  getListRunFunc(uc),
	}
}

func getListRunFunc(uc interfaces.UseCase) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		var err error
		var tasks *[]string

		if len(args) == 0 {
			tasks, err = uc.GetAll()
		} else {
			s := args[0]
			tasks, err = uc.GetByStatus(s)
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		for _, t := range *tasks {
			fmt.Println(t)
		}
	}
}
