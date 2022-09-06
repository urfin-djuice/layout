package example

import (
	"fmt"
	"github.eapteka.ru/example/internal/repository"
)

// Service Сервис Example
type Service struct {
	repo repository.Example
}

// New Создать новый сервис Example
func New(repo repository.Example) *Service {
	return &Service{
		repo: repo,
	}
}

// SomeBusinessProcess Некая бизнес логика сервиса Example
func (s *Service) SomeBusinessProcess() error {
	model, err := s.repo.GetOne()
	if err != nil {
		return err
	}
	fmt.Println(model)
	return nil
}

// OtherBusinessProcess Другая бизнес логика сервиса Example
func (s *Service) OtherBusinessProcess() error {
	models, err := s.repo.GetList()
	if err != nil {
		return err
	}
	fmt.Println(models)
	return nil
}
