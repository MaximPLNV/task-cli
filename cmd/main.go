package main

import (
	"task-cli/internal/application/usecases"
	cmd "task-cli/internal/delivery"
	"task-cli/internal/domain/interfaces"
	"task-cli/internal/infrastracture/repository"
)

var FILE_NAME string = "task.json"

func main() {
	var repo interfaces.Repository = repository.NewJsonRepository(FILE_NAME)
	var uc interfaces.UseCase = usecases.NewTaskUseCase(repo)
	cmd.Execute(uc)
}
