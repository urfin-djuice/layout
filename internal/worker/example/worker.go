package example

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/urfin-djuice/layout/internal/service"
	"github.com/urfin-djuice/layout/pkg/logger"
)

const (
	name = "ExampleWorker"
)

type Worker struct {
	service service.Example
	log     *logger.Logger
}

func New(service service.Example) *Worker {
	return &Worker{
		service: service,
		log:     logger.New(logger.LevelDebug),
	}
}

// Run Выполнение приложения
func (w *Worker) Run(_ context.Context) error {
	fmt.Println("Start worker example")

	if err := w.service.SomeBusinessProcess(); err != nil {
		return errors.Wrap(err, "Some business error")
	}

	if err := w.service.OtherBusinessProcess(); err != nil {
		return errors.Wrap(err, "Other business error")
	}

	return nil
}

func (w *Worker) Close(_ context.Context, err error) error {
	// Тут ожидание завершения и корректное завершение воркера
	fmt.Println("Close worker example")
	return err
}
