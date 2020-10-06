package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// var testFlag = flag.String("DEBUG", "1", "xxdslfkj")
var loc = flag.String("location", "World", "Name of location to greet")

func TestPingRoute(t *testing.T) {
	fmt.Printf("================================================\n")
	fmt.Printf("===> 1: %+v\n", *loc)
	fmt.Printf("================================================\n")
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}

func TestParksRoute(t *testing.T) {
	router := setupRouter()
	var createJson = []byte(`{
		"name": "aaa",
		"description": "bbb",
		"location": "ccc"}`)

	var updateJson = []byte(`{
			"name": "aaaaaa",
			"description": "bbb",
			"location": "ccc"}`)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/parks", bytes.NewBuffer(createJson))
	router.ServeHTTP(w, req)
	fmt.Printf("+++> CREATE: %+v\n", w.Body.String())
	assert.Equal(t, 201, w.Code)

	req, _ = http.NewRequest("PUT", "/parks", bytes.NewBuffer(updateJson))
	router.ServeHTTP(w, req)
	fmt.Printf("+++> UPDATE: %+v\n", w.Body.String())
	assert.Equal(t, 201, w.Code)

	req, _ = http.NewRequest("GET", "/parks", nil)
	router.ServeHTTP(w, req)
	fmt.Printf("+++> LIST: %+v\n", w.Body.String())
	assert.Equal(t, 201, w.Code)
	// fails
	// assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}
