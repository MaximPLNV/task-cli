package repository

import (
	"fmt"
	"task-cli/internal/entities"
)

func NewJsonRepository() *JsonRepository {
	return &JsonRepository{}
}

type JsonRepository struct{}

func (jr *JsonRepository) GetAll() (*[]entities.Task, error) {
	fmt.Println("GetAll Repo")

	return nil, nil
}

func (jr *JsonRepository) GetByStatus(st string) (*[]entities.Task, error) {
	fmt.Println("GetByStatus Repo")

	return nil, nil
}

func (jr *JsonRepository) Add(t *entities.Task) error {
	fmt.Println("Add Repo")

	return nil
}

func (jr *JsonRepository) Update(t *entities.Task) error {
	fmt.Println("Update Repo")

	return nil
}

func (js *JsonRepository) Delete(id int) error {
	fmt.Println("Delete Repo")

	return nil
}
