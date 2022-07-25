package stack

import (
	"runtime"
)

// Frame represents a function call and it's metadata.
type Frame struct {
	Function string `json:"function,omitempty"`
	Module   string `json:"module,omitempty"`
	Filename string `json:"filename,omitempty"`
	Lineno   int    `json:"lineno,omitempty"`
}

// NewFrame assembles a stacktrace frame out of runtime.Frame.
func NewFrame(f runtime.Frame) Frame {
	pkg, function := splitQualifiedFunctionName(f.Function)

	return Frame{
		Filename: f.File,
		Lineno:   f.Line,
		Module:   pkg,
		Function: function,
	}
}
