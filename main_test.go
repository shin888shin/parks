package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shin888shin/parks/domain/parks"
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

	// CREATE /parks //////////////////////////////////////////////////////////
	req, _ := http.NewRequest("POST", "/parks", bytes.NewBuffer(createJson))
	router.ServeHTTP(w, req)
	fmt.Printf("+++> CREATE: %+v\n", w.Body.String())
	assert.Equal(t, 201, w.Code)

	// UPDATE /parks/1 ////////////////////////////////////////////////////////
	req, _ = http.NewRequest("PUT", "/parks/1", bytes.NewBuffer(updateJson))
	router.ServeHTTP(w, req)
	fmt.Printf("+++> UPDATE: %+v\n", w.Body.String())
	assert.Equal(t, 201, w.Code)

	// LIST all parks /////////////////////////////////////////////////////////
	req, _ = http.NewRequest("GET", "/parks", nil)
	router.ServeHTTP(w, req)
	fmt.Printf("+++> LIST: %+v <+++\n", w.Body.String())
	assert.Equal(t, 201, w.Code)

	// GET parks/1 ////////////////////////////////////////////////////////////
	var park parks.Park
	req, _ = http.NewRequest("GET", "/parks/1", nil)
	router.ServeHTTP(w, req)
	fmt.Printf("+++> GET: %+v <+++\n", w.Body.String())
	assert.Equal(t, 201, w.Code)

	err := json.NewDecoder(w.Body).Decode(&park)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// spew.Dump(park)
	assert.Equal(t, int64(1), park.ID)
}
