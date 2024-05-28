package httputils_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"modules/v2/common/httputils"
	"modules/v2/common/httputils/codes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(httputils.ErrorMiddleware())

	return r
}

func TestSuccessResponse(t *testing.T) {
	r := setupRouter()
	res := httputils.NewSuccessBaseResponse(nil)

	// Set table
	tests := []struct {
		description  string
		fn           func(*gin.Context)
		expectedCode int
		path         string
	}{
		{
			// Ok Success
			description: "OKSuccess",
			fn: func(c *gin.Context) {
				res.OKSuccess(c)
			},
			expectedCode: 200,
			path:         "/test/ok",
		},
		{
			// Created Success
			description: "CreatedSuccess",
			fn: func(c *gin.Context) {
				res.CreatedSuccess(c)
			},
			expectedCode: 201,
			path:         "/test/created",
		},
	}
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tc.path, nil)

			// Specific handler function and serve
			r.GET(tc.path, tc.fn)
			r.ServeHTTP(w, req)

			if w.Code != tc.expectedCode {
				t.Errorf("Expected %d, got %d", tc.expectedCode, w.Code)
			}
		})
	}
}

func TestFailed(t *testing.T) {
	r := setupRouter()

	// Set error based data and response
	netErr := httputils.NewNetError(codes.CreateFailed, nil)
	res := httputils.NewErrorBaseResponse(netErr)

	tests := []struct {
		description  string
		fn           func(*gin.Context)
		expectedCode int
		expectedMsg  string
	}{
		{
			description: "Failed",
			fn: func(c *gin.Context) {
				res.Failed(c)
			},
			expectedCode: 500,
			expectedMsg:  "The creation of the requested resource failed.",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test/failed", nil)

			// Specific handler function and serve
			r.GET("/test/failed", tc.fn)
			r.ServeHTTP(w, req)

			// Check Status code
			if w.Code != tc.expectedCode {
				t.Errorf("Expected %d, got %d", tc.expectedCode, w.Code)
			}

			var baseRes httputils.BaseResponse
			resErr := json.Unmarshal(w.Body.Bytes(), &baseRes)
			if resErr != nil {
				t.Fatalf("Failed to unmarshal response")
			}

			// Check Message
			errData := baseRes.UnmarshalErrorData()
			if errData.Message != tc.expectedMsg {
				t.Errorf("Expected %s, got %s", tc.expectedMsg, errData.Message)
			}
		})
	}

}
