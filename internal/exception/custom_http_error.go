package exception

import "fmt"

type CustomHttpError struct {
	message string
	code    int
}

func NewCustomHttpError(code int, message string) *CustomHttpError {
	return &CustomHttpError{
		message: message,
		code:    code,
	}
}

func (err *CustomHttpError) Code() int {
	return err.code
}

func (err *CustomHttpError) Message() string {
	return err.message
}

func (err *CustomHttpError) Error() string {
	return fmt.Sprintf("CustomHttpError: %s", err.message)
}
