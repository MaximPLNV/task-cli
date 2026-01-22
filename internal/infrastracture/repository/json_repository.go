package repository

import "task-cli/internal/entities"

type JsonRepository struct{}

func (jr *JsonRepository) GetAll() ([]*entities.Task, error) {
	return nil, nil
}

func (jr *JsonRepository) GetByStatus(st string) ([]*entities.Task, error) {
	return nil, nil
}

func (jr *JsonRepository) Add(t *entities.Task) error {
	return nil
}

func (jr *JsonRepository) Update(t *entities.Task) error {
	return nil
}

func (js *JsonRepository) Delete(id int) error {
	return nil
}
