package loop

import (
	"context"
	"fmt"
	"time"
)

type Loop struct {
	iterations int64
}

func New(start int) *Loop {
	return &Loop{
		iterations: int64(start),
	}
}

func (l *Loop) Run(ctx context.Context) error {
	fmt.Println("Start worker loop")

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(1 * time.Second):
			l.iterations++
			fmt.Println("Iteration", l.iterations)
		}
	}
}

func (l *Loop) Close(_ context.Context, err error) error {
	fmt.Println("Close worker loop")
	l.iterations = 0
	return err
}
