package handlers

import (
	"encoding/json"
)

func errToJSON(err error) string {
	jsonErr, err := json.Marshal(map[string]string{"error": err.Error()})
	if err != nil {
		return ""
	}
	return string(jsonErr)
}
