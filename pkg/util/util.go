package util

import (
	"encoding/json"
	"net/http"
	"strconv"
	appError "todo-api-golang/internal/error"

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

func GetIntId(r *http.Request, uriParam string) (uint, error) {
	vars := mux.Vars(r)
	// convert the id into an integer and return
	id, err := strconv.ParseUint(vars[uriParam], 10, 64)
	if err != nil {
		return 0, appError.ErrInvalidId
	}
	return uint(id), nil
}

func ReadRequestBody[T any](r *http.Request, data *T) error {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return err
	}
	return nil
}
