package usecases

import (
	"encoding/json"
	"errors"
	"task-cli/internal/domain/entities"
	"task-cli/internal/domain/interfaces"
	"time"
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
		Status:      "ToDo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return uc.repo.Add(task)
}

func (uc *TaskUseCase) Update(id int, desc string) error {
	t, err := uc.repo.GetById(id)

	if err != nil {
		return err
	}

	t.Description = desc
	t.UpdatedAt = time.Now()

	return uc.repo.Update(t)
}

func (uc *TaskUseCase) UpdateStatus(id int, st string) error {
	t, err := uc.repo.GetById(id)

	if err != nil {
		return err
	}

	t.Status = st
	t.UpdatedAt = time.Now()

	return uc.repo.Update(t)
}

func (uc *TaskUseCase) Delete(id int) error {
	return uc.repo.Delete(id)
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
