// ParseRequestBody parses the request body based on the Content-Type header.
package minima

import (
	"encoding/json"
	"net/http"
)

func ParseRequestBody(r *http.Request) (map[string]interface{}, error) {
	if r.Body == nil {
		return nil, nil
	}
	defer r.Body.Close()

	switch r.Header.Get("Content-Type") {
	case "application/x-www-form-urlencoded":
		return parseFormData(r)
	case "application/json":
		return parseJSONData(r)
	default:
		return nil, nil // Ignore other content types
	}
}

func parseFormData(r *http.Request) (map[string]interface{}, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	data := make(map[string]interface{}, len(r.Form))
	for k, v := range r.Form {
		data[k] = v
	}

	return data, nil
}

func parseJSONData(r *http.Request) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}