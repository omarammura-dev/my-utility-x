package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRegisterRoutes(t *testing.T) {

	r := RegisterRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/url", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"message": "Hello,World!"}`, w.Body.String())
}
