package util

import (
	"encoding/json"
	"net/http"
	appError "todo-api-golang/pkg/error"

	"github.com/gorilla/mux"
)

func WriteResponse[T any](rw http.ResponseWriter, code int, data *T) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(code)
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		panic(err)
	}
}
func WriteError(rw http.ResponseWriter, error *appError.ErrorResponse) error {
	rw.Header().Add("Content-Type", "application/problem+json")
	rw.WriteHeader(error.Code)
	if err := json.NewEncoder(rw).Encode(error); err != nil {
		return err
	}
	return nil
}

func GetUriParam(r *http.Request, uriParam string) string {
	vars := mux.Vars(r)
	return vars[uriParam]
}

func ReadRequestBody[T any](r *http.Request, data *T) error {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return err
	}
	return nil
}
