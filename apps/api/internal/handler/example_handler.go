package handler

import (
	"net/http"

	"github.com/figoalfarqi/navalerp/pkg/response"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, "Hello from apipml", map[string]string{"hello": "world"}, nil)
}
