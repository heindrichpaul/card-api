package apiutilities

import (
	"log"
	"net/http"
	"net/http/httputil"

	jsoniter "github.com/json-iterator/go"
	"github.com/twinj/uuid"
)

type ApiError struct {
	ID      string
	Code    string
	Message string
}

//Marshal marshals a pointer to a Deck into a byte array for transmission.
func (z *ApiError) marshal() ([]byte, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(z)
}

func NewAPIError(message, code string) *ApiError {
	e := &ApiError{
		ID:      uuid.NewV4().String(),
		Code:    code,
		Message: message,
	}
	return e
}

func (z *ApiError) LogRequest(r *http.Request) {
	request, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println(string(request))
	}
}
