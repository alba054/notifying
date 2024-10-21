package request

type PostNotificationMessagePayload struct {
	Id      string
	Topic   string
	Message string `json:"message"`
}
