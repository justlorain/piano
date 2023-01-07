package core

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// PianoKey play the piano with PianoKeys
type PianoKey struct {
	Request *http.Request
	Writer  http.ResponseWriter

	index    int
	Params   Params
	Handlers HandlersChain
}

// NewContext will return a new context object which is piano key
func NewContext(maxParams uint16) *PianoKey {
	ps := make(Params, 0, maxParams)
	return &PianoKey{
		Params: ps,
		index:  -1,
	}
}

// Next executes the handlers on the chain
func (pk *PianoKey) Next(ctx context.Context) {
	pk.index++
	for i := pk.index; i < len(pk.Handlers); i++ {
		pk.Handlers[i](ctx, pk)
	}
}

// Query is used to match HTTP GET query params
func (pk *PianoKey) Query(key string) string {
	return pk.Request.URL.Query().Get(key)
}

// DefaultQuery is Query with default value when no match
func (pk *PianoKey) DefaultQuery(key, defaultValue string) string {
	value := pk.Query(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// PostForm is used to get HTTP POST form data
func (pk *PianoKey) PostForm(key string) string {
	return pk.Request.PostFormValue(key)
}

// DefaultPostForm is PostForm with default value when no match
func (pk *PianoKey) DefaultPostForm(key, defaultValue string) string {
	value := pk.PostForm(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func (pk *PianoKey) FormValue(key string) string {
	return pk.Request.FormValue(key)
}

func (pk *PianoKey) DefaultFormValue(key, defaultValue string) string {
	value := pk.FormValue(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// SetStatus is used to set HTTP response code
func (pk *PianoKey) SetStatus(code int) {
	pk.Writer.WriteHeader(code)
}

// SetHeader is used to set HTTP response header
func (pk *PianoKey) SetHeader(key, value string) {
	if value == "" {
		pk.Writer.Header().Del(key)
		return
	}
	pk.Writer.Header().Set(key, value)
}

func (pk *PianoKey) writeJSON(data any) error {
	pk.SetHeader("Content-Type", "application/json; charset=utf-8")
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = pk.Writer.Write(jsonBytes)
	if err != nil {
		return err
	}
	return nil
}

// JSON is used to response data in JSON form
func (pk *PianoKey) JSON(code int, data any) {
	pk.SetStatus(code)
	err := pk.writeJSON(data)
	if err != nil {
		panic(err)
	}
}

func (pk *PianoKey) writeString(format string, data ...any) error {
	pk.SetHeader("Content-Type", "text/plain; charset=utf-8")
	// Fprintf will pass the data to the writer
	_, err := fmt.Fprintf(pk.Writer, format, data...)
	return err
}

// String is used to response data in string form
func (pk *PianoKey) String(code int, format string, data ...any) {
	pk.SetStatus(code)
	err := pk.writeString(format, data...)
	if err != nil {
		panic(err)
	}
}

// refresh will reset the PianoKey as a new one
func (pk *PianoKey) refresh() {
	// TODO:
}
