package handler

import (
	"net/http"

	"github.com/figoalfarqi/apipml/pkg/response"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, "Hello from apipml", map[string]string{"hello": "world"}, nil)
}
