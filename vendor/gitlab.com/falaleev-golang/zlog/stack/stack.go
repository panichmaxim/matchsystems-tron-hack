package stack

import "runtime"

const stackLimit = 100

// GetStacktrace creates a stacktrace using runtime.Callers.
func GetStacktrace(skip int) []Frame {
	pcs := make([]uintptr, stackLimit)
	n := runtime.Callers(skip, pcs)

	if n == 0 {
		return nil
	}

	return extractFrames(pcs[:n])
}
