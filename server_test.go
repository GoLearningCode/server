package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func TestServer(t *testing.T) {
	s := Server{":3000"}
	s.GET("/", handler)
	s.POST("/post", handler)

	go s.Serve(nil)

	response, err := http.Get("http://localhost:3000")
	if err != nil {
		t.Fatal("error")
	}
	if response.StatusCode != 200 {
		t.Fatal("status not 200")
	}

	responsePost, err := http.Post("http://localhost:3000/post", "text/html", io.MultiReader())
	if err != nil {
		t.Fatal("error post")
	}
	if responsePost.StatusCode != 200 {
		t.Fatal("status post not 200")
	}
}
