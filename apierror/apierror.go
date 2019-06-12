package apierror

import (
	"log"
	"net/http"
	"net/http/httputil"

	jsoniter "github.com/json-iterator/go"
	"github.com/twinj/uuid"
)

const (
	//MarshalError is a constant error id used by all errors related to marshalling.
	MarshalError = "1"
	//NotFoundError is a constant error id used by all errors related to not being able to find the requested resource.
	NotFoundError = "2"
)

//APIError is struct that wraps the errors produced by the API.
type APIError struct {
	ID      string
	Code    string
	Message string
}

//Marshal marshals a pointer to a Deck into a byte array for transmission.
func (z *APIError) Marshal() ([]byte, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(z)
}

//NewAPIError returns a pointer to a newly initialized apierror.APIError.
func NewAPIError(message, code string) *APIError {
	e := &APIError{
		ID:      uuid.NewV4().String(),
		Code:    code,
		Message: message,
	}
	return e
}

//LogRequest logs the http.Request is used.
func (z *APIError) LogRequest(r *http.Request) {
	request, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println(string(request))
	}
}
