package server

import (
	"net/http"
)

// Defining the Server type
type Server struct {
	Address string
}

// Serve function to initialize the server
func (s Server) Serve(handler http.Handler) {
	http.ListenAndServe(s.Address, handler)
}

// http methods handler functions

// get method
func (s Server) GET(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
	http.HandleFunc(route, httpHandler)
}

// POST method
func (s Server) POST(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
	http.HandleFunc(route, httpHandler)
}

// PUT method
func (s Server) PUT(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
	http.HandleFunc(route, httpHandler)
}

// PATCH method
func (s Server) PATCH(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
	http.HandleFunc(route, httpHandler)
}

// DELETE method
func (s Server) DELETE(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
	http.HandleFunc(route, httpHandler)
}

// OPTIONS method
func (s Server) OPTIONS(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "OPTIONS" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
	http.HandleFunc(route, httpHandler)
}

// HEAD method
func (s Server) HEAD(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "HEAD" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
	http.HandleFunc(route, httpHandler)
}
