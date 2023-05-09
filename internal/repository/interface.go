package repository

import (
	"github.com/urfin-djuice/layout/internal/domain"
)

//go:generate mockgen -source $GOFILE -destination=./mock/mock.go

// Example Интерфейс репозитория Example
type Example interface {
	GetOne() (*domain.Example, error)
	GetList() ([]domain.Example, error)
}
