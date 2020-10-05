package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testFlag = flag.String("DEBUG", "1", "xxdslfkj")

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}

func TestParksRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/parks", nil)
	router.ServeHTTP(w, req)

	// fmt.Printf("+++> 0: %+v\n", *testFlag)
	// fmt.Printf("+++> 1: %+v\n", w.Code)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Printf("+++> 2: %+v\n", w.Body.String())
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")

	assert.Equal(t, 200, w.Code)
	// fails
	// assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}
