package zlog

import "gitlab.com/falaleev-golang/zlog/stack"

const stackSkip = 4

// MarshalStack implements errors stack trace marshaling.
// zerolog.ErrorStackMarshaler = MarshalStack.
func MarshalStack(err error) interface{} {
	return stack.GetStacktrace(stackSkip)
}
