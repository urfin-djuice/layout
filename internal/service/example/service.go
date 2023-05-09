package example

import (
	"fmt"

	"github.com/urfin-djuice/layout/internal/repository"
)

// Service Сервис Example
type Service struct {
	repo   repository.Example
	env    string
	domain string
}

// New Создать новый сервис Example
func New(env, domain string, repo repository.Example) *Service {
	return &Service{
		repo:   repo,
		domain: domain,
		env:    env,
	}
}

// SomeBusinessProcess Некая бизнес логика сервиса Example
func (s *Service) SomeBusinessProcess() error {
	s.printEnv()
	model, err := s.repo.GetOne()
	if err != nil {
		return err
	}
	fmt.Println(model)
	return nil
}

// OtherBusinessProcess Другая бизнес логика сервиса Example
func (s *Service) OtherBusinessProcess() error {
	s.printEnv()
	models, err := s.repo.GetList()
	if err != nil {
		return err
	}
	fmt.Println(models)
	return nil
}

func (s *Service) printEnv() {
	fmt.Println(s.env, s.domain)
}
