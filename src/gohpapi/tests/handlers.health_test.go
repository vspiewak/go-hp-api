package tests

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
  "gohpapi/handlers"
)

// Test that health check returns 200 & status up
func TestGetHealth(t *testing.T) {

	assert := assert.New(t)

	r := getRouter(true)

	r.GET("/", handlers.GetHealth)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) {

		p, err := ioutil.ReadAll(w.Body)

		assert.Nil(err)
		assert.Equal(w.Code, http.StatusOK, "Should return StatusOK")
		assert.JSONEq(string(p), `{"status": "UP"}`, "Should return a JSON with a status field and UP value")

	})

}
