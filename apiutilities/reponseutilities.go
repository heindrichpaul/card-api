package apiutilities

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/heindrichpaul/card-api/apierror"
)

//Marshalable is an interface that represents all structs that can be marshalled into JSON.
type Marshalable interface {
	Marshal() ([]byte, error)
}

//ResponseType is a struct that wraps the response the API should send.
type ResponseType struct {
	StatusCode int
	JSON       string
}

func verifyAndMarshal(m Marshalable) (response ResponseType) {
	marshalJSON, err := m.Marshal()
	if err != nil {
		e := apierror.NewAPIError(fmt.Sprintf("Could not marshal %s", reflect.TypeOf(m).Name()), apierror.MarshalError)
		return verifyAndMarshal(e)
	}
	response.JSON = string(marshalJSON)
	if isError(m) {
		response.StatusCode = 500
	} else {
		response.StatusCode = 200
	}

	return
}

func isError(m Marshalable) bool {
	typeOfMarshalable := reflect.TypeOf(m)
	var nameOfMarshalable string
	if nameOfMarshalable = string(typeOfMarshalable.Kind()); nameOfMarshalable == "Ptr" {
		nameOfMarshalable = string(typeOfMarshalable.Elem().Kind())
	}
	return strings.Contains(nameOfMarshalable, "ApiError")
}

//HandleResponse accepts the Marshalable paramenter and writes it to the given http.ResponseWriter
func HandleResponse(w http.ResponseWriter, m Marshalable) {
	w.Header().Set("Content-Type", "application/json")
	response := verifyAndMarshal(m)
	w.WriteHeader(response.StatusCode)
	fmt.Fprintf(w, response.JSON)
}
