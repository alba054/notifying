package exception

import "fmt"

type NotFoundError struct {
	message string
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		message: message,
	}
}

func (err *NotFoundError) Code() int {
	return 404
}

func (err *NotFoundError) Message() string {
	return err.message
}

func (err *NotFoundError) Error() string {
	return fmt.Sprintf("NotFoundError: %s", err.message)
}
