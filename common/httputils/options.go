package httputils

import (
	"sync"
)

var (
	opts netErrorOptions
)

// netErrorOptions is meta data about service
type netErrorOptions struct {
	// tag is the service name
	tag string
}

// SetTag sets the service name
func SetTag(tag string) {
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()

	opts.tag = tag
}
