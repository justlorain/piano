# PIANO (dev-version)

> Piano will respond to you.

![piano](image/piano.png)

Piano is a simple HTTP framework.

## Quick Start

### Hello World

[example](example/helloworld/main.go)

```go
package main

import (
	"context"
	"net/http"
	"piano/core"
)

func main() {
	p := core.Default()
	p.GET("/hello", func(ctx context.Context, pk *core.PianoKey) {
		pk.String(http.StatusOK, "piano")
	})
	_ = p.Play()
}
```

PIANO is a subproject of the [BINARY WEB ECOLOGY](https://github.com/B1NARY-GR0UP)