package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	direktori, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(direktori))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response, _ := io.ReadAll(recorder.Result().Body)
	// fmt.Println(string(response))

	assert.Equal(t, "Hallo Text", string(response))
}

func TestServeFileGood(t *testing.T) {
	router := httprouter.New()
	direktori, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(direktori))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/files/good.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response, _ := io.ReadAll(recorder.Result().Body)
	// fmt.Println(string(response))

	assert.Equal(t, "good", string(response))
}
