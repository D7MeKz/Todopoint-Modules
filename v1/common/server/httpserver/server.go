package httpserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func NewServer(routes *gin.Engine, opt ...Option) (*http.Server, error) {
	var o options
	for _, opt := range opt {
		err := opt(&o)
		if err != nil {
			return nil, err
		}
	}

	addr := o.getAddr()
	srv := http.Server{
		Addr:           addr,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &srv, nil
}
