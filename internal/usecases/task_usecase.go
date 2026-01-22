package usecases

import (
	"task-cli/internal/entities"
	"task-cli/internal/interfaces"
)

type TaskUseCase struct {
	repo interfaces.Repository
}

func (uc *TaskUseCase) GetAll() ([]*entities.Task, error) {
	return nil, nil
}

func (uc *TaskUseCase) GetByStatus(st string) ([]*entities.Task, error) {
	return nil, nil
}

func (uc *TaskUseCase) Add(desc string) error {
	return nil
}

func (uc *TaskUseCase) Update(id int, desc string) error {
	return nil
}

func (uc *TaskUseCase) UpdateStatus(st string) error {
	return nil
}

func (uc *TaskUseCase) Delete(id int) error {
	return nil
}
