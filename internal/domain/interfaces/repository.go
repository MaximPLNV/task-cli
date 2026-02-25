package interfaces

import "task-cli/internal/domain/entities"

type Repository interface {
	GetById(int) (*entities.Task, error)
	GetAll() (*[]entities.Task, error)
	GetByStatus(string) (*[]entities.Task, error)
	Add(*entities.Task) error
	Update(*entities.Task) error
	Delete(int) error
}
