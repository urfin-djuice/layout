package service

//go:generate mockgen -source $GOFILE -destination=./mock/mock.go

type Example interface {
	SomeBusinessProcess() error
	OtherBusinessProcess() error
}
