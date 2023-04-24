package main 

import (
	"sync"
)

// hold backend servers
type Backend struct {
	URL *url.URLAlive
	Alive bool
	mux sync.Mutex
	ReverseProxy *httputil.ReverseProxy
}

// Track the backends
type ServerPool struct {
	backends []*Backend
	currentIndex int
}