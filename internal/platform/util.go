package platform

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func WriteResponse[T any](rw http.ResponseWriter, code int, data T) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(code)
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		panic(err)
	}
}

func GetIntId(r *http.Request, uriParam string) uint {
	vars := mux.Vars(r)
	// convert the id into an integer and return
	id, err := strconv.ParseUint(vars[uriParam], 10, 64)
	if err != nil {
		panic(err)
	}
	return uint(id)
}

func ReadRequestBody[T any](r *http.Request, data *T) error {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return err
	}
	return nil
}
