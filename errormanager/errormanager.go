package errormanager

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	jsoniter "github.com/json-iterator/go"
	"github.com/twinj/uuid"
)

type APIError struct {
	Id      string
	Code    string
	Message string
}

//UnmarshalDeck unmarshals a byte array into a pointer to a Deck for internal use.
func UnmarshalDeck(data []byte) (*APIError, error) {
	var r APIError
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(data, &r)
	return &r, err
}

//Marshal marshals a pointer to a Deck into a byte array for transmission.
func (z *APIError) Marshal() ([]byte, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(z)
}

func NewAPIError(message, code string) *APIError {
	e := &APIError{
		Id:      uuid.NewV4().String(),
		Code:    code,
		Message: message,
	}
	return e
}

func NewAPIErrorAndLog(message, code string, r *http.Request) *APIError {
	e := &APIError{
		Id:      uuid.NewV4().String(),
		Code:    code,
		Message: message,
	}
	request, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(string(request))
	}

	return e
}
