package gosafe

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type GoGroup struct {
	g *errgroup.Group
}

func (g *GoGroup) Go(f func() error) {
	g.g.Go(safe(f))
}

func (g *GoGroup) Wait() error {
	return g.g.Wait()
}

func (g *GoGroup) SetLimit(limit int) {
	g.g.SetLimit(limit)
}

func (g *GoGroup) TryGo(f func() error) bool {
	return g.g.TryGo(safe(f))
}

func GoGroupWithContext(ctx context.Context) (*GoGroup, context.Context) {
	g, ctx := errgroup.WithContext(ctx)
	return &GoGroup{g: g}, ctx
}
