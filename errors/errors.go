package Errors

import "fmt"

// Creates an error based on two strings
func NewError(errorType string, errorString string) error {
	return fmt.Errorf("%s : %s", errorType, errorString)
}
