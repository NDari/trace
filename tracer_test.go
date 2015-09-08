package trace

import (
	"bytes"
	"testing"
)

func testNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := N(&buf)
	if tracer == nil {
		t.Error("Return from new should not be nil")
	} else {
		tracer.Trace("Hello trace package.")
		if buf.String() != "Hello trace package.\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
}