package apiutilities

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type Marshalable interface {
	Marshal() ([]byte, error)
}

type ResponseType struct {
	StatusCode int
	JSON       string
}

func verifyAndMarshal(m Marshalable) (response ResponseType) {
	marshalJSON, err := m.Marshal()
	if err != nil {
		errorString := fmt.Sprintf("Could not marshal %s", reflect.TypeOf(m).Name())
		e := NewAPIError(errorString, "1")
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

func HandleResponse(w http.ResponseWriter, r *http.Request, m Marshalable) {
	w.Header().Set("Content-Type", "application/json")
	response := verifyAndMarshal(m)
	w.WriteHeader(response.StatusCode)
	fmt.Fprintf(w, response.JSON)
}
