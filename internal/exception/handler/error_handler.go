package errorhandler

import (
	"alba054/kartjis-notify/internal/exception"
	webresponse "alba054/kartjis-notify/internal/model/web"
	"alba054/kartjis-notify/shared"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func UseErrorHandler(router *httprouter.Router) {
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		err, _ := i.(exception.HttpError)

		if err == nil {
			internalErr := i.(error)
			log.Println(internalErr)
			err = exception.NewCustomHttpError(500, internalErr.Error())
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(err.Code())
		encoder := json.NewEncoder(w)
		response := webresponse.ErrorResponse{
			Error: webresponse.ErrorBody{
				Code:    "E400",
				Message: err.Message(),
			},
			ApiResponse: webresponse.ApiResponse{
				Status: webresponse.Failed,
				Data:   nil,
			},
		}

		err_ := encoder.Encode(response)
		shared.ThrowError(err_)
	}
}
