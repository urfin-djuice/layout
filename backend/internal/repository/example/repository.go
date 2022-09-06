package example

import "github.eapteka.ru/example/internal/model"

type Repository struct {
	// Тут необходимые для работы репозитория переменные (конекты к БД, логеры, контекты)
}

// New Создать новый репозиторий Example
func New() *Repository {
	return &Repository{}
}

// GetOne Получить одну запись модели Example
func (r *Repository) GetOne() (*model.Example, error) {
	return &model.Example{
		Field1: "example string",
		Field2: 1,
	}, nil
}

// GetList Получить список записей модели Example
func (r *Repository) GetList() ([]model.Example, error) {
	return []model.Example{
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
