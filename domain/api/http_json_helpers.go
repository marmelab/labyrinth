package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

// parseJsonBody reads a JSON entity from the body.
func parseJsonBody(r *http.Request, entity interface{}) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return errors.New("invalid content-type, expected application/json")
	}

	err := json.NewDecoder(r.Body).Decode(entity)
	if err != nil {
		return err
	}

	return nil
}

// writeJsonResponse writes a JSON response to the HTTP writer.
func writeJsonResponse(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encodedBody, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Write(encodedBody)
	return nil
}
