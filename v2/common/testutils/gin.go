package testutils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"net/url"
)

func GetTestGinContext() *gin.Context {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	testUrl := &url.URL{
		Path: "/test",
	}

	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    testUrl,
	}
	return ctx
}
