package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParsePayload(r *http.Request, payload json) error {
	if r.Body == nil {
		return fmt.Errorf("request body empty")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteResponse(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteResponse(w, status, map[string]string{"error": err.Error()})
}
