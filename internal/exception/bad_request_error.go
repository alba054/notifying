package exception

import "fmt"

type BadRequestError struct {
	message string
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{
		message: message,
	}
}

func (err *BadRequestError) Code() int {
	return 400
}

func (err *BadRequestError) Message() string {
	return err.message
}

func (err *BadRequestError) Error() string {
	return fmt.Sprintf("BadRequestError: %s", err.message)
}
