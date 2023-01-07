package main

import (
	"context"
	"github.com/B1NARY-GR0UP/piano/core"
	"net/http"
)

func main() {
	p := core.Default()
	p.GET("/hello", func(ctx context.Context, pk *core.PianoKey) {
		pk.String(http.StatusOK, "piano")
	})
	p.Play()
}
