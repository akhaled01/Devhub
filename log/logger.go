package log

import (
	"fmt"
	"log/slog"
)

// logs an info to the terminal
func InfoConsoleLog(msg string, args ...any) {
	slog.Info(msg, args...)
}

// logs an error to the terminal
func ErrorConsoleLog(msg string, args ...any) {
	slog.Error(msg, args...)
}

// logs a warning to the console
func WarnConsoleLog(msg string, args ...any) {
	slog.Warn(msg, args...)
}

func PrintErrorTrace(err error) {
	type causer interface {
		Cause() error
	}

	fmt.Println("Error Trace:")

	for err != nil {
		fmt.Printf("- %v\n", err)

		// Check if the error implements causer interface
		if c, ok := err.(causer); ok {
			err = c.Cause()
		} else {
			break
		}
	}
}
