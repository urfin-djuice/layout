package eggs

import (
	"context"
	"errors"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type Worker interface {
	Run(ctx context.Context) error
	Close(ctx context.Context, err error) error
}

type Eggs struct {
	group    *errgroup.Group
	ctx      context.Context
	procList []Worker
	stop     context.CancelFunc
}

func New() *Eggs {
	var result Eggs
	var ctx context.Context

	ctx, result.stop = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	result.group, result.ctx = errgroup.WithContext(ctx)
	return &result
}

func (e *Eggs) Add(w Worker) {
	e.procList = append(e.procList, w)
}

func (e *Eggs) Run() error {
	for ix := range e.procList {
		w := e.procList[ix]
		e.group.Go(func() error {
			ctx2, stop := context.WithCancel(e.ctx)
			defer stop()
			err := w.Run(ctx2)
			return w.Close(ctx2, err)
		})
	}

	err := e.group.Wait()
	if errors.Is(err, context.Canceled) || err == nil {
		err = nil
	}
	e.stop()
	time.Sleep(5 * time.Second)

	return err
}
