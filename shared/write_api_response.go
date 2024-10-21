package shared

import (
	webresponse "alba054/kartjis-notify/internal/model/web"
	"encoding/json"
	"net/http"
)

func WriteApiResponse(w http.ResponseWriter, code int, status webresponse.ResponseStatus, data interface{}) {
	response := webresponse.ApiResponse{
		Status: status,
		Data:   data,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	ThrowError(err)
}
