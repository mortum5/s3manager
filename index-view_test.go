package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexViewHandler(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		expectedStatusCode   int
		expectedBodyContains string
	}{
		"success": {
			expectedStatusCode:   http.StatusPermanentRedirect,
			expectedBodyContains: "Redirect",
		},
	}

	for _, tc := range tests {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(err)

		rr := httptest.NewRecorder()
		handler := IndexViewHandler()

		handler.ServeHTTP(rr, req)

		assert.Equal(tc.expectedStatusCode, rr.Code)
		assert.Contains(rr.Body.String(), tc.expectedBodyContains)
	}
}