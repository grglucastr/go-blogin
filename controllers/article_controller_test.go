
package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/grglucastr/go-blogin/router"
)

func TestInfoRequest(t *testing.T) {
	// setup a new request
	req, err := http.NewRequest("GET", "/hellos", nil)

	if err != nil {
		t.Fatal(err)
	}

	newRecorder := httptest.NewRecorder()

	router.Routes().ServeHTTP(newRecorder, req)

	expectedStatusCode := 200
	statusCode := newRecorder.Result().StatusCode

	if expectedStatusCode != statusCode {
		t.Errorf("TestInfoRequest() test return an unexpected result: got %v want %v", statusCode, expectedStatusCode)
		t.Error(newRecorder.Body)
	}

}
