package main

import (
	"healthcheck/config"
	"healthcheck/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	config := config.New()
	app := router.New(config)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/healthcheck", nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

func TestPostLineAuthorization(t *testing.T) {
	config := config.New()
	app := router.New(config)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/line/authorization", nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

func TestGetLineAuthorization(t *testing.T) {
	config := config.New()
	app := router.New(config)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/line/authorization", nil)

	q := req.URL.Query()
	q.Add("code", "ubuKmaLgf8JlXjJG8YR9")
	q.Add("state", "MTU4MzY5MDE4NjczNDE4ODAwMA%3D%3D")
	req.URL.RawQuery = q.Encode()

	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}
