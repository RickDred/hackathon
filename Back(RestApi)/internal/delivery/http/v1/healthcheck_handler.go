package v1

import (
	"net/http"
)

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := map[string]interface{}{
		"status": "available",
		"system_info": map[string]string{
			"version": "1.0",
		},
	}

	WriteJSON(w, http.StatusOK, env, nil)
}
