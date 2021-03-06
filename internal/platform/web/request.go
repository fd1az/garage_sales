package web

import (
	"encoding/json"
	"net/http"
)

//Decoder looks for a JSON document in the request body and unmarshals it into val
func Decoder(r *http.Request, val interface{}) error {

	if err := json.NewDecoder(r.Body).Decode(&val); err != nil {

		return NewRequestError(err, http.StatusBadRequest)
	}
	return nil
}
