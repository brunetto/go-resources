package main

import (
	"errors"
	"fmt"
)

// sentinel error
const awsomeError = "my awsome error"

var err1 = errors.New(awsomeError)

// typed error
type customError struct{ err error }

func (ce customError) Error() string { return ce.err.Error() }

func main() {
	// sentinel error
	err := func() error { return err1 }()
	fmt.Println("errors.Is checks that the error we got is or wraps the variable used as sentinel error:", errors.Is(err, err1))

	var tErr customError
	fmt.Println("errors.As checks that the error we got is or wraps the same type of the error we want to check:", errors.As(err, &tErr))

	// type error
	err = func() error { return customError{err1} }()
	fmt.Println("errors.As checks that the error we got is or wraps "+
		"the same type of the error we want to check:", errors.As(err, &tErr))
	fmt.Println("errors.Is checks that the error we got is or wraps "+
		"the variable used as sentinel error:", errors.Is(err, err1))
	fmt.Println("errors.Is checks that the error we got is or wraps "+
		"the variable used as sentinel error, and now that we \"asserted\" the error type this works:", errors.Is(tErr.err, err1))
}
