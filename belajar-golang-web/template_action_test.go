package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", Page{
		Title: "Template Action If",
	})
}

func TemplateActionOperator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/operator.gohtml"))
	t.ExecuteTemplate(w, "operator.gohtml", map[string]interface{}{
		"Title":      "Template Action Operator",
		"FinalValue": 50,
	})
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Title":   "Template Action Operator",
		"Hobbies": []string{},
	})
}

func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))
	t.ExecuteTemplate(w, "with.gohtml", map[string]interface{}{
		"Title": "Template Action Operator",
		"Name":  "Vini",
		"Address": map[string]interface{}{
			"Street": "Jl. Aku Saja",
			"City":   "Ayam",
		},
	})
}
func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	// TemplateActionIf(recorder, request)
	// TemplateActionOperator(recorder, request)
	// TemplateActionRange(recorder, request)
	TemplateActionWith(recorder, request)

	response, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(response))

}
