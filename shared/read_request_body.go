package shared

import (
	"alba054/kartjis-notify/internal/exception"
	"encoding/json"

	"io"
)

func ReadRequestBody(r io.Reader, payload interface{}) {
	decoder := json.NewDecoder(r)

	err := decoder.Decode(&payload)

	if err != nil && err.Error() == "EOF" {
		panic(exception.NewBadRequestError("provide request body"))
	}

	ThrowError(err)
}
