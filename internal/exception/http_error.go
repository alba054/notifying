package exception

type HttpError interface {
	Code() int
	Message() string
}
