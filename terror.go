package terror

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

// Terror represents an error with stack trace details
type Terror struct {
	message  string
	cause    error
	file     string
	line     int
	funcName string
}

// New creates a new TreeError with the current file and line information
func New(message string, cause error) error {
	pc, file, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	return &Terror{
		message:  message,
		cause:    cause,
		file:     file,
		line:     line,
		funcName: funcName,
	}
}

// Error implements the error interface for TreeError
func (te *Terror) Error() string {
	return te.message
}

// Unwrap returns the cause of the TreeError
func (te *Terror) Unwrap() error {
	return te.cause
}

// FormatTree formats the error and its causes in a tree structure with better visuals
func FormatTree(err error) string {
	var builder strings.Builder
	printErrorTree(&builder, err, 0)
	return builder.String()
}

func printErrorTree(builder *strings.Builder, err error, level int) {
	if err == nil {
		return
	}

	// Create tree formatting symbols, but skip the prefix for the first line
	prefix := ""
	if level > 0 {
		prefix = strings.Repeat("    ", level-1) + "└── "
	}

	// Get file, line, and function information for Terror
	var file string
	var line int
	var funcName string
	if te, ok := err.(*Terror); ok {
		file = te.file
		line = te.line
		funcName = te.funcName
	}

	// Print the current error with details
	builder.WriteString(fmt.Sprintf("%s%s (%s:%d in %s)\n", prefix, err.Error(), file, line, funcName))

	// Recursively print the cause of the error if any
	if unwrappedErr := errors.Unwrap(err); unwrappedErr != nil {
		printErrorTree(builder, unwrappedErr, level+1)
	}
}
