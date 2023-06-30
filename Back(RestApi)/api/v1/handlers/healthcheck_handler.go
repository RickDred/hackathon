package handlers

import (
	"encoding/json"
	"net/http"
)

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := map[string]interface{}{
		"status": "available",
		"system_info": map[string]string{
			"version": "1.0",
		},
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(env)
}
