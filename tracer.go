package trace

import (
	"fmt"
	"io"
)

// tracer type has an io.Writer field called out which is where we will write the trace output to,
// io.Writer means that the user can decide where the tracing output will be written.
type tracer struct {
	out io.Writer
}

// nilTracer struct has defined a Trace method that does nothing,
// and a call to the Off method will create a new nilTracer struct and return it.
type nilTracer struct{}

// Tracer is the interface that describes an object capable of
// tracing events throughout code.
type Tracer interface {
	Trace(...interface{})
}

// Trace method formats and write the trace details to the out writer.
func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// New creates a new instance of a Tracer that will write the output to
// the specific io.Writer.
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

func (t *nilTracer) Trace(a ...interface{}) {}

// Off creates a Tracer that will ignore calls to Trace.
func Off() Tracer {
	return &nilTracer{}
}
