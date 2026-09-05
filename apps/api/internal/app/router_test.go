package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/figoalfarqi/apipml/config"
)

func TestSetupRouterRegistersRoutesWithoutConflict(t *testing.T) {
	router := SetupRouter(&config.Config{JWTKey: "test-secret"})
	request := httptest.NewRequest(http.MethodGet, "/api/v1/hello", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("GET /api/v1/hello returned %d", recorder.Code)
	}
}
