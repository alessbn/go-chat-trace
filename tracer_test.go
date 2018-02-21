package trace

import (
	"bytes"
	"testing"
)

// TestNew tests the tracing behaviour.
func TestNew(t *testing.T) {
	// buf captures the output of the tracer this will ensure that the string in the buffer
	// matches the expected value.
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from New should not be nil.")
	} else {
		tracer.Trace("Hello trace package.")
		if buf.String() != "Hello trace package./n" {
			t.Error("Trace should not writen '%s'.", buf.String())
		}
	}
}

// TestOff gets a silent tracer before making a call to Trace
// to ensure the code doesn't panic.
func TestOff(t *testing.T) {
	// var silentTracer Tracer = Off()
	silentTracer := Off()
	silentTracer.Trace("something")
}
