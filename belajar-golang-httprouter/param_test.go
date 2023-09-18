package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestParam(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		text := "Product " + params.ByName("id")
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/product/3", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response, _ := io.ReadAll(recorder.Result().Body)
	// fmt.Println(string(response))

	assert.Equal(t, "Product 3", string(response))
}
