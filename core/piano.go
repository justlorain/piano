package core

import (
	"context"
)

// Piano will respond to you.
type Piano struct {
	*Engine
}

type HandlerFunc func(ctx context.Context, pk *PianoKey)

// HandlersChain is the slice of HandlerFunc
type HandlersChain []HandlerFunc

func New(opts ...Option) *Piano {
	options := NewOptions(opts...)
	p := &Piano{
		Engine: NewEngine(options),
	}
	return p
}

func Default(opts ...Option) *Piano {
	p := New(opts...)
	// TODO: use recovery middleware
	// TODO: use logger middleware
	return p
}
