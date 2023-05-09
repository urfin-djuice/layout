package example

import (
	"github.com/urfin-djuice/layout/internal/domain"
)

type Repository struct {
	// Тут необходимые для работы репозитория переменные (конекты к БД, логеры, контекты)
}

// New Создать новый репозиторий Example
func New() *Repository {
	return &Repository{}
}

// GetOne Получить одну запись модели Example
func (r *Repository) GetOne() (*domain.Example, error) {
	return &domain.Example{
		Field1: "example string",
		Field2: 1,
	}, nil
}

// GetList Получить список записей модели Example
func (r *Repository) GetList() ([]domain.Example, error) {
	return []domain.Example{
		{
			Field1: "string 1",
			Field2: 1,
		},
		{
			Field1: "string 2",
			Field2: 1,
		},
		{
			Field1: "string 3",
			Field2: 1,
		},
	}, nil
}
