package main

import (
	"log"
	"net/http"
	"time"
)

// Logger — middleware to log all incoming requests.
type Logger struct {
	handler http.Handler
}

// NewLogger — constructor for a Logger struct.
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handler: handlerToWrap}
}

// ServeHTTP — method to satisfy http.Handler interface.
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}
