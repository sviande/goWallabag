package goWallabag

import (
	"testing"
)

var hasBeenClose = false

type Closer struct {
	HasBeenClose bool
}

func (c Closer) Close() error {
	hasBeenClose = true
	return nil
}

func TestDeferClose(t *testing.T) {
	closer := Closer{}
	deferClose(closer)
	if !hasBeenClose {
		t.Error("DeferClose doesn't call close on closer")
	}
}
