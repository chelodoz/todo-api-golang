package util

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBody[T any](r *http.Request, data *T) error {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return err
	}
	return nil
}
