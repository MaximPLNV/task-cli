package interfaces

type UseCase interface {
	GetAll() (*[]string, error)
	GetByStatus(string) (*[]string, error)
	Add(string) error
	Update(int, string) error
	UpdateStatus(int, string) error
	Delete(int) error
}
