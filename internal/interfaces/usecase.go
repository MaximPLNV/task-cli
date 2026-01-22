package interfaces

import "task-cli/internal/entities"

type UseCase interface {
	GetAll() ([]*entities.Task, error)
	GetByStatus(string) ([]*entities.Task, error)
	Add(string) error
	Update(int, string) error
	UpdateStatus(string) error
	Delete(int) error
}
