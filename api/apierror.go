package api

import (
	"log"
	"net/http"
	"net/http/httputil"

	jsoniter "github.com/json-iterator/go"
	"github.com/twinj/uuid"
)

type apiError struct {
	ID      string
	Code    string
	Message string
}

//Marshal marshals a pointer to a Deck into a byte array for transmission.
func (z *apiError) marshal() ([]byte, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(z)
}

func newAPIError(message, code string) *apiError {
	e := &apiError{
		ID:      uuid.NewV4().String(),
		Code:    code,
		Message: message,
	}
	return e
}

func (z *apiError) logRequest(r *http.Request) {
	request, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println(string(request))
	}
}
