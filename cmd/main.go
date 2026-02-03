/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	cmd "task-cli/internal/infrastracture/delivery"
	"task-cli/internal/infrastracture/repository"
	"task-cli/internal/interfaces"
	"task-cli/internal/usecases"
)

func main() {
	var repo interfaces.Repository = repository.NewJsonRepository()
	var uc interfaces.UseCase = usecases.NewTaskUseCase(repo)
	cmd.Execute(uc)
}
