package main 

import (
	"sync"
	"net/http"
	"net/url"
	"net/http/httputil"
	"log"
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

// Relay requests through ReverseProxy
u, _ := url.Parse("http://localhost:8080")
reverse_proxy := httputil.NewSingleHostReverseProxy(u)

// Initialize the server pool
http.HandlerFunc(reverse_proxy.ServeHTTP)