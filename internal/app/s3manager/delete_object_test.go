package s3manager_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mastertinner/s3manager/internal/app/s3manager"
	"github.com/stretchr/testify/assert"
)

func TestDeleteObjectHandler(t *testing.T) {
	cases := map[string]struct {
		removeObjectFunc     func(string, string) error
		expectedStatusCode   int
		expectedBodyContains string
	}{
		"deletes an existing object": {
			removeObjectFunc: func(bucketName string, objectName string) error {
				return nil
			},
			expectedStatusCode:   http.StatusNoContent,
			expectedBodyContains: "",
		},
		"returns error if there is an S3 error": {
			removeObjectFunc: func(bucketName string, objectName string) error {
				return errors.New("mocked S3 error")
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedBodyContains: http.StatusText(http.StatusInternalServerError),
		},
	}

	for tcID, tc := range cases {
		t.Run(tcID, func(t *testing.T) {
			assert := assert.New(t)

			s3 := &S3Mock{
				RemoveObjectFunc: tc.removeObjectFunc,
			}

			req, err := http.NewRequest(http.MethodDelete, "/api/buckets/bucketName/objects/objectName", nil)
			assert.NoError(err, tcID)

			rr := httptest.NewRecorder()
			handler := s3manager.DeleteObjectHandler(s3)

			handler.ServeHTTP(rr, req)

			assert.Equal(tc.expectedStatusCode, rr.Code, tcID)
			assert.Contains(rr.Body.String(), tc.expectedBodyContains, tcID)
		})
	}
}