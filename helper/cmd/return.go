package cmd

import "fmt"

type ExitCodeError struct {
	Message string
	Code    int
}

func (e ExitCodeError) Error() string {
	return fmt.Sprintf("exiting: %s (%d)", e.Message, e.Code)
}
