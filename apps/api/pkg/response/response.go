package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, message string, data interface{}, errors map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    status,
		"status":  http.StatusText(status),
		"message": message,
		"data":    data,
		"errors":  errors,
	})
}
