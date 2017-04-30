package tests

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {

	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())

}

// Helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	return r
}

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder)) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service
	r.ServeHTTP(w, req)

	// process the above request
	f(w)

}
