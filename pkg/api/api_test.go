package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/health", nil)

	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(HealthCheck)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Health Check returned status, %v ", rec.Code)
	}
	want := "OK"
	assert.Equal(t, want, rec.Body.String())
}
