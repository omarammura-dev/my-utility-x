package testing

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"myutilityx.com/routes"
)

func TestHelloReturns200(t *testing.T) {

	r := routes.RegisterRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"message": "Hello,World!"}`, w.Body.String())
}
