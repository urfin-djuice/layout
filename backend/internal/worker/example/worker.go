package example

import (
	repo "github.eapteka.ru/example/internal/repository/example"
	serv "github.eapteka.ru/example/internal/service/example"
	"github.eapteka.ru/example/pkg/logger"
	"log"
)

const (
	name = "ExampleWorker"
)

type Worker struct {
	service *serv.Service
	log     *logger.Logger
}

func New() *Worker {
	return &Worker{
		service: serv.New(repo.New()),
		log:     logger.New(logger.LevelDebug),
	}
}

// Run Выполнение приложения
func (w *Worker) Run() {
	if err := w.init(); err != nil {
		log.Fatal("Fail to init ", name, " - ", err.Error())
	}

	if err := w.service.SomeBusinessProcess(); err != nil {
		w.log.Error("Some error", err)
	}

	if err := w.service.OtherBusinessProcess(); err != nil {
		w.log.Error("Other error", err)
	}

	w.closer()
}

func (w *Worker) closer() {
	// Тут ожидание завершения и корректное завершение воркера
}

func (w *Worker) init() error {
	// Тут инициализация приложения (конфиги, контексты, конекты к БД и т.д.)
	return nil
}
