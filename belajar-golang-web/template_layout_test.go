package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	//ini menggunakna nama alias aau define nama di file gohtmlnya
	t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Template layout",
		"Name":  "Vini",
	})
}

func TestLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	response, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(response))

}
