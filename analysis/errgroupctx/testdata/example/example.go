package example

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func Ok() error {
	g, ctx2 := errgroup.WithContext(context.Background())

	g.Go(func() error {
		select {
		case <-ctx2.Done():
			return ctx2.Err()
		default:
			return fmt.Errorf("some error")
		}
	})

	return g.Wait()
}

func OkEmpty() error {
	group, _ := errgroup.WithContext(context.TODO())
	group.Go(func() error { return nil })
	return group.Wait()
}

func ErrOverwriteCtx() error {
	var g *errgroup.Group
	var ctx context.Context

	g, ctx = errgroup.WithContext(context.Background()) // want "context from errgroup.WithContext should not be named 'ctx'"

	g.Go(func() error {
		return nil
	})

	_ = ctx
	return g.Wait()
}

func ErrNewCtxOnAssignment() error {
	g, ctx := errgroup.WithContext(context.Background()) // want "context from errgroup.WithContext should not be named 'ctx'"

	g.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return fmt.Errorf("some error")
		}
	})

	return g.Wait()
}
