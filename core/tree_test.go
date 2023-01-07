package core

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestMatchKind(t *testing.T) {
	cKind := matchKind("name")
	pKind := matchKind(":name")
	wKind := matchKind("*name")
	fmt.Println(cKind)
	fmt.Println(pKind)
	fmt.Println(wKind)
}

func TestInsert(t *testing.T) {
	tree := &tree{
		method: http.MethodGet,
		root:   nil,
	}
	tree.insert("/auth/hello", HandlersChain{
		func(ctx context.Context, pk *PianoKey) {
			pk.String(200, "wow")
		},
	})
	tree.insert("/auth/world", HandlersChain{
		func(ctx context.Context, pk *PianoKey) {
			pk.JSON(200, &M{
				"code": 1,
				"msg":  "success",
			})
		},
	})
}
