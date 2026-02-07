package interfaces

import "task-cli/internal/entities"

type Repository interface {
	GetAll() (*[]entities.Task, error)
	GetByStatus(string) (*[]entities.Task, error)
	Add(*entities.Task) error
	Update(*entities.Task) error
	Delete(int) error
}
