package webresponse

type ResponseStatus string

const (
	Success ResponseStatus = "SUCCESS"
	Failed  ResponseStatus = "FAILED"
)

type ApiResponse struct {
	Status ResponseStatus `json:"status"`
	Data   interface{}    `json:"data"`
}

type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"messagge"`
}

type ErrorResponse struct {
	ApiResponse
	Error ErrorBody `json:"error"`
}
