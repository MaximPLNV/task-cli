package usecases

import (
	"task-cli/internal/entities"
	"task-cli/internal/interfaces"
)

func NewTaskUseCase(repo interfaces.Repository) *TaskUseCase {
	return &TaskUseCase{
		repo: repo,
	}
}

type TaskUseCase struct {
	repo interfaces.Repository
}

func (uc *TaskUseCase) GetAll() ([]*string, error) {
	result, err := uc.repo.GetAll()
	//json.Marshal(result)
	return nil, err
}

func (uc *TaskUseCase) GetByStatus(st string) ([]*string, error) {
	result, err := uc.repo.GetByStatus(st)
	//json.Marshal(result)
	return nil, err
}

func (uc *TaskUseCase) Add(desc string) error {
	task := &entities.Task{
		Description: desc,
		Status:      "todo",
	}
	err := uc.repo.Add(task)

	return err
}

func (uc *TaskUseCase) Update(id int, desc string) error {
	task := &entities.Task{
		Id:          id,
		Description: desc,
	}
	err := uc.repo.Update(task)

	return err
}

func (uc *TaskUseCase) UpdateStatus(id int, st string) error {
	task := &entities.Task{
		Id:     id,
		Status: st,
	}
	err := uc.repo.Update(task)

	return err
}

func (uc *TaskUseCase) Delete(id int) error {
	err := uc.repo.Delete(id)

	return err
}
