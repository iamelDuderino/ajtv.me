package main

import (
	"io"
	"net/http"
)

// custom response writer to handle gzip
type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
