// Package web provides utilities for constructing HTTP responses.
package web

import (
	"net/http"
)

// ContentType sets the Content-Type header.
func ContentType(w http.ResponseWriter, ct string) {
	Header(w, "Content-Type", ct)
}

// Header sets a header key to a single value.
func Header(w http.ResponseWriter, key, value string) {
	w.Header().Set(key, value)
}
