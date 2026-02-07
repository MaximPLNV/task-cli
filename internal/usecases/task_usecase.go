package usecases

import (
	"encoding/json"
	"errors"
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

func (uc *TaskUseCase) GetAll() (*[]string, error) {
	tasks, err := uc.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return uc.parseTasksToStr(tasks)
}

func (uc *TaskUseCase) GetByStatus(st string) (*[]string, error) {
	tasks, err := uc.repo.GetByStatus(st)

	if err != nil {
		return nil, err
	}

	return uc.parseTasksToStr(tasks)
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

func (uc *TaskUseCase) parseTasksToStr(tasks *[]entities.Task) (*[]string, error) {
	if tasks == nil {
		return nil, errors.New("There is no tasks")
	}

	result := make([]string, len(*tasks))

	for i, t := range *tasks {
		b, jsonErr := json.Marshal(t)

		if jsonErr != nil {
			return nil, jsonErr
		}

		result[i] = string(b)
	}
	return &result, nil
}
