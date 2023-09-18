package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	//Proses Parsing
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	first_name := r.PostForm.Get("first_name")
	last_name := r.PostForm.Get("last_name")

	//cara instant tanpa membuat parsing manual
	// first_name := r.PostFormValue("first_name")

	fmt.Fprintf(w, "Hello %s %s", first_name, last_name)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Vini&last_name=Jumatul")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}
