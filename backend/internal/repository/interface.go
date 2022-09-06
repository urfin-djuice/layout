package repository

// Тут интерфейсы для генерации моков и написания юнит тестов для сервисов
//go:generate mockgen -source $GOFILE -destination=./mock/mock.go

import "github.eapteka.ru/example/internal/model"

// Example Интерфейс репозитория Example
type Example interface {
	GetOne() (*model.Example, error)
	GetList() ([]model.Example, error)
}
