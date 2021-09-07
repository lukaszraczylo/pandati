package pandati

import (
	"fmt"
	"runtime"
)

// CheckForError takes err and additional message as parameter,
// Returning pre-formatted error message with stack trace if err is not nil
func CheckForError(err error, msg ...string) (toReturn string) {
	if err != nil {
		toReturn = fmt.Sprintf("ERROR: %s (%s) | %s", msg, err, Trace())
	}
	return
}

// Trace returns trace information about executing function together line number and file name
// It's useful for debugging purposes
func Trace() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("TRACE: %s:%d %s\n", frame.File, frame.Line, frame.Function)
}
